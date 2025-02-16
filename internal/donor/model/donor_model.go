package model

import (
	"time"

	"github.com/charitan-go/profile-server/pkg/proto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Donor struct {
	id         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ReadableId uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName  string    `gorm:"type:varchar(255);not null;unique"`
	LastName   string    `gorm:"type:varchar(255);not null;unique"`
	Address    string    `gorm:"type:varchar(255)"`
	createdAt  time.Time
	updatedAt  time.Time
}

func (d *Donor) BeforeCreate(db *gorm.DB) (err error) {
	if d.id == uuid.Nil {
		d.id = uuid.New()
	}

	if d.ReadableId == uuid.Nil {
		d.ReadableId = uuid.New()
	}

	return nil
}

func NewDonor(reqDto *proto.CreateDonorProfileRequestDto) *Donor {
	return &Donor{
		FirstName: reqDto.FirstName,
		LastName:  reqDto.LastName,
		Address:   reqDto.Address,
	}
}

func (d *Donor) ToCreateDonorProfileResponseDto() *proto.CreateDonorProfileResponseDto {
	result := &proto.CreateDonorProfileResponseDto{
		ProfileReadableId: d.ReadableId.String(),
	}
	return result
}
