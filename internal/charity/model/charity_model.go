package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Charity struct {
	id               uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ReadableId       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	OrganizationName string    `gorm:"type:varchar(255);not null;unique"`
	TaxCode          string    `gorm:"type:varchar(255);not null;unique"`
	Address          string    `gorm:"type:varchar(255);not null"`
	createdAt        time.Time
	updatedAt        time.Time
}

func (c *Charity) BeforeCreate(db *gorm.DB) (err error) {
	if c.id == uuid.Nil {
		c.id = uuid.New()
	}

	if c.ReadableId == uuid.Nil {
		c.ReadableId = uuid.New()
	}

	return nil
}

// func NewCharity(reqDto *proto.CreateCharityProfileRequestDto) *Charity {
// 	return &Charity{
// 		FirstName: reqDto.FirstName,
// 		LastName:  reqDto.LastName,
// 		Address:   reqDto.Address,
// 	}
// }
