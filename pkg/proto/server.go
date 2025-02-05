package proto

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type ProfileProtoServer struct {
	UnimplementedProfileProtoServiceServer
}

func SetupGrpcServiceServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	profileProtoServer := grpc.NewServer()
	RegisterProfileProtoServiceServer(profileProtoServer, &ProfileProtoServer{})
	log.Println("GRPC server listening on :50051")
	if err := profileProtoServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(profileProtoServer, healthServer)

	healthServer.SetServingStatus("profile-server-grpc", grpc_health_v1.HealthCheckResponse_SERVING)
}

func (s *ProfileProtoServer) CreateDonorProfile(
	ctx context.Context,
	req *CreateDonorProfileRequestDto,
) (*CreateDonorProfileResponseDto, error) {
	log.Println(req.FirstName)
	return &CreateDonorProfileResponseDto{ProfileReadableId: "32312432"}, nil
}
