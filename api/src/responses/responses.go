package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a response in Json to the request
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Fatal(err)
	}
}

// Retorn an error in format JSON.
func Error(w http.ResponseWriter, code int, err error) {

	JSON(w, code, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})

}
