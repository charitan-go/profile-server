package rest

import (
	"fmt"
	"log"
	"os"

	"github.com/charitan-go/profile-server/rest/api"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/labstack/echo/v4"
)

type RestServer struct {
	echo *echo.Echo
	api  *api.Api
}

func NewEcho() *echo.Echo {
	return echo.New()
}

func NewRestServer(echo *echo.Echo, api *api.Api) *RestServer {
	return &RestServer{echo, api}
}

func (s *RestServer) setupRouting() {
	s.echo.GET("/health", s.api.HealthCheck)

	// Endpoint for profile
}

func (s *RestServer) setupServiceRegistry() {
	log.Println("Start for service discovery")

	config := consulapi.DefaultConfig()
	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println("Cannot connect with service registry", err)
	}

	address := os.Getenv("ADDRESS")

	// Register REST service (Echo app)
	restServiceId := fmt.Sprintf("%s-rest", address)
	restRegistration := &consulapi.AgentServiceRegistration{
		Name:    restServiceId,
		ID:      restServiceId,
		Address: address,
		Port:    8090,
		Tags:    []string{"rest"},
		Check: &consulapi.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:8090/health", address),
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	err = consul.Agent().ServiceRegister(restRegistration)
	if err != nil {
		log.Fatalf("Failed to register REST service with Consul: %v", err)
	} else {
		log.Println("Setup service registry for grpc service ok")
	}
}

func (s *RestServer) setup() {
	// Setup rest api routingj
	s.setupRouting()

	// Setup service registry
	s.setupServiceRegistry()
}

func (s *RestServer) Run() {
	log.Println("Start run rest server")

	s.setup()

	s.echo.Start(":8090")
	log.Println("Server started at http://localhost:8090")
}
