package repository

type TrackerRepository interface {
	Register()
}

type trackerRepository struct {
}

func (tr *trackerRepository) Register() {}

func NewTrackerRepository() TrackerRepository {
	return &trackerRepository{}
}
