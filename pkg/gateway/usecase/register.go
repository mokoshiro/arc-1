package usecase

import (
	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
)

type RegisterUsecase interface {
	Register(req *model.RegisterRequest) error
}

type registerUsecase struct {
	metadataRepo repository.MetadataRepository
	trackerRepo  repository.TrackerRepository
}

func (ru *registerUsecase) Register(req *model.RegisterRequest) error {
	return nil
}

func NewRegisterUsecase(metaRepo repository.MetadataRepository, trackerRepo repository.TrackerRepository) RegisterUsecase {
	return &registerUsecase{
		metadataRepo: metaRepo,
		trackerRepo:  trackerRepo,
	}
}
