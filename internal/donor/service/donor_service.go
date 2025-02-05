package service

import "github.com/charitan-go/profile-server/internal/donor/repository"

type DonorService interface{}

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
