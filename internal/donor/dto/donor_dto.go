package dto

import "github.com/google/uuid"

type DonorDto struct {
	ReableId  uuid.UUID `json:"readableId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   string    `json:"address"`
}
