package usecase

import "github.com/Bo0km4n/arc/pkg/gateway/domain/repository"

type RegisterUsecase interface {
	Register()
}

type registerUsecase struct {
	metadataRepo repository.MetadataRepository
	trackerRepo  repository.TrackerRepository
}

func (ru *registerUsecase) Register() {

}

func NewRegisterUsecase(metaRepo repository.MetadataRepository, trackerRepo repository.TrackerRepository) RegisterUsecase {
	return &registerUsecase{
		metadataRepo: metaRepo,
		trackerRepo:  trackerRepo,
	}
}
