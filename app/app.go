package app

import (
	"github.com/charitan-go/profile-server/grpc"
	"github.com/charitan-go/profile-server/pkg/database"
	"github.com/charitan-go/profile-server/rest"
)

func Run() {
	// Connect to db
	database.SetupDatabase()

	// Run proto server
	go grpc.Run()

	// Run rest server
	go rest.Run()

	// Prevent main from exit
	select {}
}
