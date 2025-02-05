package donor

import (
	"github.com/charitan-go/profile-server/internal/donor/handler"
	"github.com/charitan-go/profile-server/internal/donor/repository"
	"github.com/charitan-go/profile-server/internal/donor/service"
	"go.uber.org/fx"
)

var DonorModule = fx.Module("donor",
	fx.Provide(
		repository.NewDonorRepository,
		service.NewDonorService,
		handler.NewDonorHandler,
	),
)
