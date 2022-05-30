package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/helpers/json_response"
)

var errInternalServer error = errors.New("internal server error")

func internalServerErrorHandler(w http.ResponseWriter, errorPrefix string, err error) bool {
	if errors.Is(err, errInternalServer) {
		fmt.Printf("%s: %s", errorPrefix, err)
		err = json_response.InternalServerError(w, err)
		if err != nil {
			fmt.Printf("%s: %s", errorPrefix, err)
		}
		return true
	}
	return false
}

func getParamFromRequestFloat64(r *http.Request, paramName string) float64 {
	value, err := strconv.ParseFloat(chi.URLParam(r, paramName), 64)
	if err != nil {
		panic(fmt.Errorf("%w: failed to parse {%s}: %s", errInternalServer, paramName, err))
	}
	return value
}

func getParamFromRequestInt64(r *http.Request, paramName string) int64 {
	value, err := strconv.ParseInt(chi.URLParam(r, paramName), 10, 64)
	if err != nil {
		panic(fmt.Errorf("%w: failed to parse {%s}: %s", errInternalServer, paramName, err))
	}
	return value
}

//TODO: move CatchErrorWrapper to another file

type ErrorHandler func(http.ResponseWriter, string, error) bool

type CatchErrorsWrapper struct {
	Prefix   string
	Handlers []ErrorHandler
}

func (cw CatchErrorsWrapper) Wrapp(f http.HandlerFunc, handlers ...ErrorHandler) http.HandlerFunc {
	cw.Handlers = handlers
	return func(w http.ResponseWriter, r *http.Request) {
		defer cw.catchErrors(w)
		f(w, r)
	}
}

func (cw *CatchErrorsWrapper) catchErrors(w http.ResponseWriter) {
	if recoverer := recover(); recoverer != nil {
		if err, ok := recoverer.(error); ok {
			for _, handler := range cw.Handlers {
				if handler(w, cw.Prefix, err) {
					return
				}
			}
		}
		panic(recoverer)
	}
}
