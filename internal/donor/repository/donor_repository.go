package repository

import (
	"github.com/charitan-go/profile-server/internal/donor/model"
	"github.com/charitan-go/profile-server/pkg/database"
	"gorm.io/gorm"
)

type DonorRepository interface {
	Save(*model.Donor) (*model.Donor, error)

	FindOneByReadableId(readableId string) (*model.Donor, error)
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

// FindOneByReadableId implements DonorRepository.
func (r *donorRepositoryImpl) FindOneByReadableId(readableId string) (*model.Donor, error) {
	var donor model.Donor

	result := r.db.Where("readableId = ?", readableId).First(&donor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &donor, nil
}
