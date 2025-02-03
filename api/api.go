package api

import (
	"net/http"

	charity "github.com/charitan-go/profile-server/domain/charity/handler"
	donor "github.com/charitan-go/profile-server/domain/donor/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	CharityHandler *charity.CharityHandler
	DonorHandler *donor.DonorHandler
}

func NewApi(charityHandler *charity.CharityHandler, donorHandler *donor.DonorHandler) *Api {
	return &Api{
		CharityHandler: charityHandler,
		DonorHandler: donorHandler,
	}
}

func (*Api) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
