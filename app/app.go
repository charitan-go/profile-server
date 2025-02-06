package app

import (
	"log"

	"github.com/charitan-go/profile-server/grpc"
	"github.com/charitan-go/profile-server/internal/charity"
	"github.com/charitan-go/profile-server/internal/donor"
	"github.com/charitan-go/profile-server/pkg/database"
	"github.com/charitan-go/profile-server/rest"
	"github.com/charitan-go/profile-server/rest/api"
	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(restSrv *rest.RestServer, grpcSrv *grpc.GrpcServer) {
	log.Println("In invoke")
	// Start REST server
	go func() {
		log.Println("In goroutine of rest")
		restSrv.Run()
	}()

	// Start gRPC server
	go func() {
		log.Println("In goroutine of grpc")
		grpcSrv.Run()
	}()
}

func Run() {
	// Connect to db
	database.SetupDatabase()

	fx.New(
		donor.DonorModule,
		charity.CharityModule,
		fx.Provide(
			rest.NewRestServer,
			rest.NewEcho,
			api.NewApi,
			grpc.NewGrpcServer,
		),
		fx.Invoke(runServers),
	).Run()
}
