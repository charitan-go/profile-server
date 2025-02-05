package charity

import (
	"github.com/charitan-go/profile-server/internal/charity/handler"
	"go.uber.org/fx"
)

var CharityModule = fx.Module("charity",
	fx.Provide(
		handler.NewCharityHandler,
	),
)
