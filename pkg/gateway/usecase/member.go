package usecase

import (
	"context"
	"time"

	"errors"

	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
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
	res := &model.GetMemberByRadiusResponse{
		Members: []*model.Member{},
	}
	var unit proto.GetMemberByRadiusRequest_Unit

	switch req.Unit {
	case "km":
		unit = proto.GetMemberByRadiusRequest_KM
	case "m":
		unit = proto.GetMemberByRadiusRequest_M
	case "mi":
		unit = proto.GetMemberByRadiusRequest_MI
	case "ft":
		unit = proto.GetMemberByRadiusRequest_FT
	default:
		unit = proto.GetMemberByRadiusRequest_KM
	}

	trackerRes, err := muc.trackerRepo.GetMemberByRadius(context.Background(), &proto.GetMemberByRadiusRequest{
		Longitude: req.Location.Longitude,
		Latitude:  req.Location.Latitude,
		Radius:    req.Radius,
		Unit:      unit,
	})
	if err != nil {
		return nil, err
	}
	res.Members = make([]*model.Member, len(trackerRes.Members))
	for i, v := range trackerRes.Members {
		res.Members[i] = &model.Member{
			Location: &model.Location{
				Latitude:  v.Latitude,
				Longitude: v.Longitude,
			},
			ID: v.PeerId,
		}
	}
	return res, nil
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
