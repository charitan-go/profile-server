package dto

import "github.com/google/uuid"

type DonorDto struct {
	ReableId  uuid.UUID `json:"readable_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Address   string    `json:"address"`
}
