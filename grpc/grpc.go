package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/charitan-go/profile-server/internal/donor"
	donorservice "github.com/charitan-go/profile-server/internal/donor/service"
	"github.com/charitan-go/profile-server/pkg/proto"
	consulapi "github.com/hashicorp/consul/api"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type GrpcServer struct {
	proto.UnimplementedProfileGrpcServiceServer
	donorSvc donorservice.DonorService
}

func newGrpcServer(donorSvc donorservice.DonorService) *GrpcServer {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	profileGrpcServer := &GrpcServer{}
	proto.RegisterProfileGrpcServiceServer(grpcServer, profileGrpcServer)
	profileGrpcServer.donorSvc = donorSvc

	address := os.Getenv("SERVICE_ID")
	grpcServiceName := fmt.Sprintf("%s-grpc", address)
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus(grpcServiceName, grpc_health_v1.HealthCheckResponse_SERVING)

	log.Println("GRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return profileGrpcServer
}

func (*GrpcServer) setupServiceRegistry() {
	log.Println("Start for service discovery")

	config := consulapi.DefaultConfig()
	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("Cannot connect with service registry", err)
	}

	address := os.Getenv("ADDRESS")

	grpcServiceId := fmt.Sprintf("%s-grpc", address)
	grpcRegistration := &consulapi.AgentServiceRegistration{
		ID:      grpcServiceId,
		Name:    grpcServiceId,
		Address: address,
		Port:    50051,
		Tags:    []string{"grpc"},
		Check: &consulapi.AgentServiceCheck{
			GRPC:     fmt.Sprintf("%v:%d", address, 50051),
			Interval: "10s",
			Timeout:  "5s",
			// DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = consul.Agent().ServiceRegister(grpcRegistration)
	if err != nil {
		log.Fatalf("Failed to register gRPC service with Consul: %v", err)
	}
}

func (s *GrpcServer) CreateDonorProfile(
	ctx context.Context,
	req *proto.CreateDonorProfileRequestDto,
) (*proto.CreateDonorProfileResponseDto, error) {
	log.Println(req.FirstName)
	return &proto.CreateDonorProfileResponseDto{ProfileReadableId: "32312432"}, nil
}

func Run() {
	log.Println("Start run proto server")

	fx.New(
		donor.DonorModule,

		fx.Provide(
			newGrpcServer,
		),
		fx.Invoke(func(s *GrpcServer) {
			s.setupServiceRegistry()
		}),
	).Run()
}
