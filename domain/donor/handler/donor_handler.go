package handler

type DonorHandler struct{}

func (h *DonorHandler) CheckHealth() string {
	return "OK"
}

func NewDonorHandler() *DonorHandler {
	return &DonorHandler{}
}

// func (h *DonorHandler) RegisterDonor(c echo.Context) DonorDto {
// }
