package repository

import (
	"github.com/charitan-go/profile-server/internal/donor/model"
	"github.com/charitan-go/profile-server/pkg/database"
	"gorm.io/gorm"
)

type DonorRepository interface {
	Save(*model.Donor) (*model.Donor, error)
}

type donorRepositoryImpl struct {
	db *gorm.DB
}

func NewDonorRepository() DonorRepository {
	return &donorRepositoryImpl{db: database.DB}
}

func (r *donorRepositoryImpl) Save(donorModel *model.Donor) (*model.Donor, error) {
	result := r.db.Create(donorModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return donorModel, nil
}
