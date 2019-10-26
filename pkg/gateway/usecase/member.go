package usecase

import (
	"context"
	"time"

	"errors"

	"github.com/Bo0km4n/arc/pkg/gateway/domain/model"
	"github.com/Bo0km4n/arc/pkg/gateway/domain/repository"
	metaproto "github.com/Bo0km4n/arc/pkg/metadata/api/proto"
	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	"golang.org/x/xerrors"
)

type MemberUsecase interface {
	Register(req *model.RegisterRequest) error
	GetMemberByRadius(req *model.GetMemberByRadiusRequest) (*model.GetMemberByRadiusResponse, error)
	Update(req *model.UpdateRequest) (*model.UpdateResponse, error)
	Delete(req *model.DeleteRequest) (*model.DeleteResponse, error)
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
	ctx := context.Background()
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

	trackerRes, err := muc.trackerRepo.GetMemberByRadius(ctx, &proto.GetMemberByRadiusRequest{
		Longitude: req.Location.Longitude,
		Latitude:  req.Location.Latitude,
		Radius:    req.Radius,
		Unit:      unit,
	})
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(trackerRes.Members))
	for i := range trackerRes.Members {
		ids[i] = trackerRes.Members[i].PeerId
	}
	metadataRes, err := muc.metadataRepo.GetMember(ctx, &metaproto.GetMemberRequest{
		PeerIds: ids,
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
			ID:   v.PeerId,
			Addr: metadataRes.Members[i].Addr,
		}
	}

	return res, nil
}

func (mu *memberUsecase) Update(req *model.UpdateRequest) (*model.UpdateResponse, error) {
	ctx := context.Background()
	member, err := mu.metadataRepo.GetMember(ctx, &metaproto.GetMemberRequest{
		PeerIds: []string{req.ID},
	})
	if err != nil {
		return nil, err
	}
	if len(member.Members) == 0 {
		return nil, xerrors.Errorf("Couldn't get key: %s", req.ID)
	}
	if _, err := mu.trackerRepo.Update(ctx, &proto.UpdateRequest{
		PeerId:    req.ID,
		Longitude: req.Location.Longitude,
		Latitude:  req.Location.Latitude,
	}); err != nil {
		return nil, err
	}
	return &model.UpdateResponse{}, nil
}

func (mu *memberUsecase) Delete(req *model.DeleteRequest) (*model.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	done := make(chan error, 1)

	// Get lock
	go func() {
		err := mu.lockerRepo.Lock(ctx, req.ID)
		if err != nil {
			done <- err
		}
		done <- nil
	}()

	select {
	case e := <-done:
		if e != nil {
			return nil, e
		}
		break
	case <-ctx.Done():
		return nil, errors.New("Timeout get lock of a key")
	}
	defer mu.lockerRepo.Unlock(ctx, req.ID)

	if _, err := mu.metadataRepo.Delete(ctx, &metaproto.DeleteRequest{
		PeerId: req.ID,
	}); err != nil {
		return nil, err
	}
	return &model.DeleteResponse{}, nil

	// TODO: Delete member from Tracker
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
