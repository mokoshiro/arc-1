package repository

import (
	"context"

	"github.com/Bo0km4n/arc/pkg/tracker/api/proto"
	trackerclient "github.com/Bo0km4n/arc/pkg/tracker/client"
	"golang.org/x/xerrors"
)

type TrackerRepository interface {
	Register(ctx context.Context, peerID string, latitude, longitude float64) error
}

type trackerRepository struct {
	trackerclient.Client
}

func (tr *trackerRepository) Register(ctx context.Context, peerID string, latitude, longitude float64) error {
	_, err := tr.Client.Register(ctx, &proto.RegisterRequest{
		Id:        peerID,
		Latitude:  latitude,
		Longitude: longitude,
	})
	if err != nil {
		return xerrors.Errorf("trackerRepository.Register: %w", err)
	}
	return nil
}

func NewTrackerRepository(trackerClient trackerclient.Client) TrackerRepository {
	return &trackerRepository{
		trackerClient,
	}
}
