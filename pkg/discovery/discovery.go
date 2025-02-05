package discovery

import (
	"fmt"
	"log"
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

func SetupServiceRegistry() {
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

	// Register REST service (Echo app)
	restServiceId := fmt.Sprintf("%s-rest", address)
	restRegistration := &consulapi.AgentServiceRegistration{
		Name:    restServiceId,
		ID:      restServiceId,
		Address: address,
		Port:    8090,
		Tags:    []string{"http"},
		Check: &consulapi.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:8090/health", address),
			Interval:                       "10s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	// Register both services
	err = consul.Agent().ServiceRegister(grpcRegistration)
	if err != nil {
		log.Fatalf("Failed to register gRPC service with Consul: %v", err)
	}

	err = consul.Agent().ServiceRegister(restRegistration)
	if err != nil {
		log.Fatalf("Failed to register REST service with Consul: %v", err)
	}

	log.Println("Successfully registered profile-service (gRPC + REST) with Consul")
}
