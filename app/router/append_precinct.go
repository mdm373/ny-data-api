package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const idPathParam = "id"

type exampleResponse struct {
	Message string `json:"message"`
}

func writeJSON(responseWriter http.ResponseWriter, body interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(body)
}

func getPrecinctBounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, ok := vars[idPathParam]; ok {
		writeJSON(w, exampleResponse{Message: "hello world"})
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func appendPrecinct(parent *mux.Router) {
	precinctByIDRouter := parent.PathPrefix(fmt.Sprintf("/precinct/{%s}", idPathParam)).Subrouter()
	precinctByIDRouter.HandleFunc("/bounds", getPrecinctBounds).Methods("GET")
}
