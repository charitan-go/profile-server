package repository

import (
	"github.com/charitan-go/profile-server/internal/charity/model"
	"github.com/charitan-go/profile-server/pkg/database"
	"gorm.io/gorm"
)

type CharityRepository interface {
	Save(*model.Charity) (*model.Charity, error)

	FindOneByReadableId(readableId string) (*model.Charity, error)
}

type CharityRepositoryImpl struct {
	db *gorm.DB
}

func NewCharityRepository() CharityRepository {
	return &CharityRepositoryImpl{db: database.DB}
}

func (r *CharityRepositoryImpl) Save(CharityModel *model.Charity) (*model.Charity, error) {
	result := r.db.Create(CharityModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return CharityModel, nil
}

// FindOneByReadableId implements CharityRepository.
func (r *CharityRepositoryImpl) FindOneByReadableId(readableId string) (*model.Charity, error) {
	var Charity model.Charity

	result := r.db.Where("readable_id = ?", readableId).First(&Charity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Charity, nil
}
