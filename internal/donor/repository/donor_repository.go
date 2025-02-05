package repository

import (
	"github.com/charitan-go/profile-server/pkg/database"
	"gorm.io/gorm"
)

type DonorRepository interface{}

type donorRepositoryImpl struct {
	db *gorm.DB
}

func NewDonorRepository() DonorRepository {
	return &donorRepositoryImpl{db: database.DB}
}
