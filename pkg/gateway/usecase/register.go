package usecase

import (
	"context"
	"time"

	"errors"

	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
)

type MemberUsecase interface {
	Register(req *model.RegisterRequest) error
	GetMemberByRadius(req *model.GetMemberByRadiusRequest) (*model.GetMemberByRadiusResponse, error)
}

type memberUsecase struct {
	metadataRepo repository.MetadataRepository
	trackerRepo  repository.TrackerRepository
	lockerRepo   repository.LockerRepository
}

func (ru *memberUsecase) Register(req *model.RegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	done := make(chan error, 1)

	// Async get lock of a key
	go func() {
		err := ru.lockerRepo.Lock(ctx, req.ID)
		if err != nil {
			done <- err
		}
		done <- nil
	}()

	select {
	case e := <-done:
		if e != nil {
			return e
		}
		break
	case <-ctx.Done():
		return errors.New("Timeout get lock of a key")
	}
	defer ru.lockerRepo.Unlock(ctx, req.ID)

	if err := ru.metadataRepo.Register(ctx, req.ID, req.GlobalIPAddr+":"+req.Port); err != nil {
		return err
	}

	if err := ru.trackerRepo.Register(ctx, req.ID, req.Location.Latitude, req.Location.Longitude); err != nil {
		// rollback
		return err
	}

	return nil
}

func (muc *memberUsecase) GetMemberByRadius(req *model.GetMemberByRadiusRequest) (*model.GetMemberByRadiusResponse, error) {
	return nil, nil
}

func NewMemberUsecase(
	metaRepo repository.MetadataRepository,
	trackerRepo repository.TrackerRepository,
	lockerRepo repository.LockerRepository,
) MemberUsecase {
	return &memberUsecase{
		lockerRepo:   lockerRepo,
		metadataRepo: metaRepo,
		trackerRepo:  trackerRepo,
	}
}
