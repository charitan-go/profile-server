package charity

import (
	"github.com/charitan-go/profile-server/internal/charity/handler"
	"github.com/charitan-go/profile-server/internal/charity/repository"
	"github.com/charitan-go/profile-server/internal/charity/service"
	"go.uber.org/fx"
)

var CharityModule = fx.Module("charity",
	fx.Provide(
		handler.NewCharityHandler,
		service.NewCharityService,
		repository.NewCharityRepository,
	),
)
