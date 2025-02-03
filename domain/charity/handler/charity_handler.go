package handler

type CharityHandler struct{}

func (h *CharityHandler) CheckHealth() string {
	return "OK"
}

func NewCharityHandler() *CharityHandler {
	return &CharityHandler{}
}
