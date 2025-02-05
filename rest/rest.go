package rest

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/genproto/protobuf/api"
)

// type Api struct {
// 	CharityHandler *charity.CharityHandler
// 	DonorHandler *donor.DonorHandler
// }

type Rest struct {
	echo *echo.Echo
	api  *api.Api
}
