package proto

import (
	context "context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServiceServer struct {
	UnimplementedProfileProtoServiceServer
}

var grpcServiceServer *GrpcServiceServer

func SetupGrpcServiceServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServiceServer := grpc.NewServer()
	RegisterProfileProtoServiceServer(grpcServiceServer, &GrpcServiceServer{})
	log.Println("GRPC server listening on :50051")
	if err := grpcServiceServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcServiceServer) CreateDonorProfile(
	ctx context.Context,
	req *CreateDonorProfileRequestDto,
) (*CreateDonorProfileResponseDto, error) {
	log.Println(req.FirstName)
	return &CreateDonorProfileResponseDto{ProfileReadableId: "32312432"}, nil
}
