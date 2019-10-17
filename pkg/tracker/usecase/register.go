package usecase

import (
	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/domain/repository"
)

type RegisterUsecase interface {
	Register(req *proto.RegisterRequest) error
}

type registerUsecase struct {
	rr repository.RegisterRepository
}

func (ru *registerUsecase) Register(req *proto.RegisterRequest) error {
	// preprocess request
	return ru.rr.Register("H3 Hash", req.Id, req.Longitude, req.Latitude)
}

func NewRegisterUsecase(rr repository.RegisterRepository) RegisterUsecase {
	return &registerUsecase{
		rr: rr,
	}
}
