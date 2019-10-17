package repository

import (
	trackerclient "github.com/Bo0km4n/arc/pkg/tracker/client"
)

type TrackerRepository interface {
	Register()
}

type trackerRepository struct {
	trackerclient.Client
}

func (tr *trackerRepository) Register() {

}

func NewTrackerRepository(trackerClient trackerclient.Client) TrackerRepository {
	return &trackerRepository{
		trackerClient,
	}
}
