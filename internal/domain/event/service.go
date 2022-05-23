package event

import "encoding/json"

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(*Event) error
	Update(id int64, decoder *json.Decoder) (*Event, error)
	Delete(id int64) error
	GetNearby(long, lat, distance float64) ([]Event, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return s.repo.FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return s.repo.FindOne(id)
}

func (s *service) Create(event *Event) error {
	return s.repo.Create(event)
}

func (s *service) Update(id int64, decoder *json.Decoder) (*Event, error) {
	return s.repo.Update(id, decoder)
}

func (s *service) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *service) GetNearby(long, lat, distance float64) ([]Event, error) {
	return s.repo.GetNearby(&Coords{long, lat}, distance)
}
