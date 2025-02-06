package service

import (
	"log"

	"github.com/charitan-go/profile-server/internal/donor/model"
	"github.com/charitan-go/profile-server/internal/donor/repository"
	"github.com/charitan-go/profile-server/pkg/proto"
)

type DonorService interface {
	CreateDonorProfile(*proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error)
}

type donorServiceImpl struct {
	r repository.DonorRepository
}

func NewDonorService(r repository.DonorRepository) DonorService {
	return &donorServiceImpl{r}
}

func NewExternalDonorService() DonorService {
	r := repository.NewDonorRepository()
	return &donorServiceImpl{r}
}

func (svc *donorServiceImpl) CreateDonorProfile(reqDto *proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error) {
	donorModel := model.NewDonor(reqDto)

	// Save to db
	donorModel, err := svc.r.Save(donorModel)
	if err != nil {
		log.Println("Cannot save donorModel")
		return nil, err
	}

	return donorModel.ToCreateDonorProfileResponseDto(), nil
}
