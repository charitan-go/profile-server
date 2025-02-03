package app

import (
	"fmt"

	"github.com/charitan-go/profile-server/api"
	charity "github.com/charitan-go/profile-server/domain/charity/handler"
	donor "github.com/charitan-go/profile-server/domain/donor/handler"
	"github.com/charitan-go/profile-server/pkg/database"
	"github.com/charitan-go/profile-server/pkg/discovery"
	protoserver "github.com/charitan-go/profile-server/pkg/proto/server"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type App struct {
	echo *echo.Echo

	api *api.Api
}

func newApp(echo *echo.Echo, api *api.Api) *App {
	return &App{
		echo: echo,
		api:  api,
	}
}

func newEcho() *echo.Echo {
	return echo.New()
}

func (app *App) setupRouting() {
	app.echo.GET("/health", app.api.HealthCheck)

	// Donation
	// app.echo.POST("/donation/register", app.api.DonorHandler.RegisterDonor)

	// Charity
}

func (app *App) setup() {
	// Register with service registry
	discovery.SetupServiceRegistry()

	// Connect to db
	database.SetupDatabase()

	// Setup GRPC server
	go protoserver.SetupGrpcServiceServer()

	// Setup routing
	app.setupRouting()
}

func Run() {
	fx.New(
		fx.Provide(
			newApp,
			newEcho,
			api.NewApi,
			charity.NewCharityHandler,
			donor.NewDonorHandler),
		fx.Invoke(func(app *App) {
			app.setup()

			go app.echo.Start(":8090")
			fmt.Println("Server started at http://localhost:8090")
		}),
	).Run()
}
