package service

import (
	"log"

	"github.com/charitan-go/profile-server/internal/donor/model"
	"github.com/charitan-go/profile-server/internal/donor/repository"
	"github.com/charitan-go/profile-server/pkg/proto"
)

type DonorService interface {
	// GRPC
	HandleCreateDonorProfileGrpc(*proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error)
	HandleGetDonorProfileGrpc(reqDto *proto.GetDonorProfileRequestDto) (*proto.GetDonorProfileResponseDto, error)
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

func (svc *donorServiceImpl) toCreateDonorProfileResponseDto(donor *model.Donor) *proto.CreateDonorProfileResponseDto {
	result := &proto.CreateDonorProfileResponseDto{
		ProfileReadableId: donor.ReadableId.String(),
	}
	return result
}

func (svc *donorServiceImpl) HandleCreateDonorProfileGrpc(reqDto *proto.CreateDonorProfileRequestDto) (*proto.CreateDonorProfileResponseDto, error) {
	donorModel := model.NewDonor(reqDto)

	// Save to db
	donorModel, err := svc.r.Save(donorModel)
	if err != nil {
		log.Println("Cannot save donorModel")
		return nil, err
	}

	return donorModel.ToCreateDonorProfileResponseDto(), nil
}

func (svc *donorServiceImpl) HandleGetDonorProfileGrpc(reqDto *proto.GetDonorProfileRequestDto) (*proto.GetDonorProfileResponseDto, error) {

	profileReadableId := reqDto.ProfileReadableId
	donor, err := svc.r.FindOneByReadableId(profileReadableId)
	if err != nil {
		log.Printf("Cannot find donor with donorId: %s\n", profileReadableId)
		return nil, err
	}

	resDto := &proto.GetDonorProfileResponseDto{
		FirstName: donor.FirstName,
		LastName:  donor.LastName,
		Address:   donor.Address,
	}

	return resDto, nil
}
