package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/test_server/internal/domain/event"
	"github.com/test_server/internal/helpers/json_response"
)

type EventController struct {
	service event.Service
}

func NewEventController(s event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		events, err := c.service.FindAll()
		if err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}

		err = json_response.Success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.FindAll()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}

func (c *EventController) FindOne() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get record id
		id := getParamFromRequestInt64(r, "id")
		// get record from db
		event, err := c.service.FindOne(id)
		if err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}
		// return result
		err = json_response.Success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.FindOne()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}

func (c *EventController) Create() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var event = new(event.Event)
		// decode json
		if err := json.NewDecoder(r.Body).Decode(event); err != nil {
			panic(fmt.Errorf("%w: JSON parsing error: %s", errInternalServer, err))
		}
		// create record
		if err := c.service.Create(event); err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}
		// return created event
		if err := json_response.Created(w, event); err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.Create()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}

func (c *EventController) Delete() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get record id
		id := getParamFromRequestInt64(r, "id")
		// delete record
		if err := c.service.Delete(id); err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}
		// return Success response
		if err := json_response.Success(w, nil); err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.Delete()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}

func (c *EventController) Update() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get record id
		id := getParamFromRequestInt64(r, "id")
		// make json decoder
		decoder := json.NewDecoder(r.Body)
		// update record
		event, err := c.service.Update(id, decoder)
		if err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}
		// return Success response
		if err := json_response.Success(w, event); err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.Update()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}

func (c *EventController) GetNearby() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get Long, Lat, Distance
		lon := getParamFromRequestFloat64(r, "long")
		lat := getParamFromRequestFloat64(r, "lat")
		dist := getParamFromRequestFloat64(r, "dist")
		// search record
		events, err := c.service.GetNearby(lon, lat, dist)
		if err != nil {
			panic(fmt.Errorf("%w: %s", errInternalServer, err))
		}
		// return Success response
		if err := json_response.Success(w, events); err != nil {
			fmt.Printf("EventController.GetNearby(): %s", err)
		}
	}
	return CatchErrorsWrapper{Prefix: "EventController.GetNearby()"}.Wrapp(
		handler,
		internalServerErrorHandler,
	)
}
