package app

import (
	"log"

	"github.com/charitan-go/profile-server/api"
	charity "github.com/charitan-go/profile-server/internal/charity/handler"
	donor "github.com/charitan-go/profile-server/internal/donor/handler"
	"github.com/charitan-go/profile-server/pkg/database"
	"github.com/charitan-go/profile-server/pkg/discovery"
	"github.com/charitan-go/profile-server/pkg/proto"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// type App struct {
// 	echo *echo.Echo
//
// 	api *api.Api
// }

type App struct {
	rest *rest.RestServer

	proto *proto.ProtoServer
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
	go discovery.SetupServiceRegistry()

	// Connect to db
	go database.SetupDatabase()

	// Setup GRPC server
	go proto.SetupProtoServer()

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
			log.Println("Server started at http://localhost:8090")
		}),
	).Run()
}
