package event

import (
	"encoding/json"

	"github.com/upper/db/v4"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(*Event) error
	Update(id int64, decoder *json.Decoder) (*Event, error)
	Delete(id int64) error
	GetNearby(coords *Coords, distance float64) ([]Event, error)
}

// TODO: change this value in the functions according to the pagination params
const EventsCount int64 = 10

type repository struct {
	// Some internal data
	session    db.Session
	collection db.Collection
	tableName  string
}

func NewRepository(session db.Session) Repository {
	return &repository{
		session:    session,
		collection: session.Collection("events"),
		tableName:  "events",
	}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, 0, EventsCount)
	res := r.collection.Find()
	err := res.All(&events)
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	event := Event{}
	res := r.collection.Find("id", id)
	err := res.One(&event)
	return &event, err
}

func (r *repository) Create(event *Event) error {
	return r.collection.InsertReturning(event)
}

func (r *repository) Delete(id int64) error {
	res := r.collection.Find("id", id)
	return res.Delete()
}

func (r *repository) Update(id int64, decoder *json.Decoder) (*Event, error) {
	event := &Event{}
	res := r.collection.Find("id", id)
	if err := res.One(event); err != nil {
		return nil, err
	}
	if err := decoder.Decode(event); err != nil {
		return nil, err
	}
	// if id had been changed during the parsing of the json request, we must change it back
	event.Id = id
	if err := res.Update(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *repository) GetNearby(coords *Coords, distance float64) ([]Event, error) {
	events := make([]Event, 0, EventsCount)
	res := r.collection.Find("ST_DWithin(coords, ?, ?)", coords.String(), distance)
	err := res.All(&events)
	return events, err
}
