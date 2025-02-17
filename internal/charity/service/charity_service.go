package service

import (
	"log"

	"github.com/charitan-go/profile-server/internal/charity/model"
	"github.com/charitan-go/profile-server/internal/charity/repository"
	"github.com/charitan-go/profile-server/pkg/proto"
)

type CharityService interface {
	// GRPC
	HandleCreateCharityProfileGrpc(*proto.CreateCharityProfileRequestDto) (*proto.CreateCharityProfileResponseDto, error)

	HandleGetCharityProfileGrpc(reqDto *proto.GetCharityProfileRequestDto) (*proto.GetCharityProfileResponseDto, error)
}

type charityServiceImpl struct {
	r repository.CharityRepository
}

func NewCharityService(r repository.CharityRepository) CharityService {
	return &charityServiceImpl{r}
}

func NewExternalCharityService() CharityService {
	r := repository.NewCharityRepository()
	return &charityServiceImpl{r}
}

func (svc *charityServiceImpl) toCreateCharityProfileResponseDto(charityModel *model.Charity) *proto.CreateCharityProfileResponseDto {
	return &proto.CreateCharityProfileResponseDto{
		ProfileReadableId: charityModel.ReadableId.String(),
	}
}

func (svc *charityServiceImpl) HandleCreateCharityProfileGrpc(reqDto *proto.CreateCharityProfileRequestDto) (*proto.CreateCharityProfileResponseDto, error) {
	charityModel := model.NewCharity(reqDto)

	// Save to db
	charityModel, err := svc.r.Save(charityModel)
	if err != nil {
		log.Println("Cannot save charityModel")
		return nil, err
	}

	return svc.toCreateCharityProfileResponseDto(charityModel), nil
}

func (svc *charityServiceImpl) HandleGetCharityProfileGrpc(reqDto *proto.GetCharityProfileRequestDto) (*proto.GetCharityProfileResponseDto, error) {

	profileReadableId := reqDto.ProfileReadableId
	charity, err := svc.r.FindOneByReadableId(profileReadableId)
	if err != nil {
		log.Printf("Cannot find charity with charityId: %s\n", profileReadableId)
		return nil, err
	}

	resDto := &proto.GetCharityProfileResponseDto{
		OrganizationName: charity.OrganizationName,
		TaxCode:          charity.TaxCode,
		Address:          charity.Address,
	}

	return resDto, nil
}
