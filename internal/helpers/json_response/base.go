package json_response

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(body)
}

func Created(w http.ResponseWriter, body interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	return json.NewEncoder(w).Encode(body)
}

func NoContent(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func BadRequest(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func InternalServerError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func ValidationError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func GenericError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
}

func NotFound(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	return json.NewEncoder(w).Encode(map[string]interface{}{"error": "Not Found"})
}
