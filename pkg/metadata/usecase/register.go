package usecase

import "github.com/Bo0km4n/arc/pkg/metadata/domain/repository"

type RegisterUsecase interface {
	Register(peerID, addr string) error
}

type registerUsecase struct {
	rr repository.RegisterRepository
}

func (ru *registerUsecase) Register(peerID, addr string) error {
	return ru.rr.Register(peerID, addr)
}

func NewRegisterUsecase(rr repository.RegisterRepository) RegisterUsecase {
	return &registerUsecase{
		rr: rr,
	}
}
