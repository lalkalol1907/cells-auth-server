package HttpTools

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	e := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	json.NewEncoder(w).Encode(e)
}
