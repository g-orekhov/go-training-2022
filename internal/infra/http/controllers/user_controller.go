package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/user"
	"github.com/test_server/internal/helpers/json_response"
)

type UserController struct {
	service user.Service
}

func NewUserController(s user.Service) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.service.FindAll()
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindAll(): %s", err)
			}
			return
		}

		err = json_response.Success(w, users)
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
		}
	}
}

func (c *UserController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("UserController.FindOne(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindOne(): %s", err)
			}
			return
		}
		user, err := c.service.FindOne(id)
		if err != nil {
			fmt.Printf("UserController.FindOne(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindOne(): %s", err)
			}
			return
		}

		err = json_response.Success(w, user)
		if err != nil {
			fmt.Printf("UserController.FindOne(): %s", err)
		}
	}
}

func (c *UserController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user = new(user.User)
		// decode json
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			fmt.Printf("UserController.Create(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.Create(): %s", err)
			}
			return
		}
		// create record
		if err := c.service.Create(user); err != nil {
			fmt.Printf("UserController.Create(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.Create(): %s", err)
			}
			return
		}
		// return result
		if err := json_response.Created(w, user); err != nil {
			fmt.Printf("UserController.Create(): %s", err)
		}
	}
}

func (c *UserController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("UserController.Delete(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.Delete(): %s", err)
			}
			return
		}

		if err := c.service.Delete(id); err != nil {
			fmt.Printf("UserController.Delete(): %s", err)
			err = json_response.InternalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.Delete(): %s", err)
			}
			return
		}

		err = json_response.Success(w, nil)
		if err != nil {
			fmt.Printf("UserController.Delete(): %s", err)
		}
	}
}
