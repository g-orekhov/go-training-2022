package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := c.service.FindAll()
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindAll(): %s", err)
			}
			return
		}

		err = json_response.Success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}
		event, err := c.service.FindOne(id)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		err = json_response.Success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event = new(event.Event)
		// decode json
		if err := json.NewDecoder(r.Body).Decode(event); err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}
		// create record
		if err := c.service.Create(event); err != nil {
			fmt.Printf("EventController.Create(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Create(): %s", err)
			}
			return
		}
		// return result
		if err := json_response.Created(w, event); err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}

func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}

		if err := c.service.Delete(id); err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}

		err = json_response.Success(w, nil)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get record id
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete(): %s", err)
			}
			return
		}
		// make json decoder
		decoder := json.NewDecoder(r.Body)
		// update record
		event, err := c.service.Update(id, decoder)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		// return Success response
		if err := json_response.Success(w, event); err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
}

func (c *EventController) GetNearby() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get Lon, Lat
		lon, err := strconv.ParseFloat(chi.URLParam(r, "long"), 64)
		if err != nil {
			fmt.Printf("EventController.GetNearby(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.GetNearby(): %s", err)
			}
			return
		}
		lat, err := strconv.ParseFloat(chi.URLParam(r, "lat"), 64)
		if err != nil {
			fmt.Printf("EventController.GetNearby(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.GetNearby(): %s", err)
			}
			return
		}
		dist, err := strconv.ParseFloat(chi.URLParam(r, "dist"), 64)
		if err != nil {
			fmt.Printf("EventController.GetNearby(): %s", err)
			//TODO: set default value
			dist = 10.0
		}
		fmt.Println(lon, lat, dist)

		// update record
		events, err := c.service.GetNearby(lon, lat, dist)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update(): %s", err)
			}
			return
		}
		// return Success response
		if err := json_response.Success(w, events); err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
}
