package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

//swagger:model error
type errorModel struct {
	//Tell the user what went terribly wrong
	//required: true
	Message string `json:"message"`
}

type RouteHandler func(http.ResponseWriter, *http.Request)

func RespondWithJSON(responseWriter http.ResponseWriter, body interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(body)
}

func RespondWithError(responseWriter http.ResponseWriter, body string) {
	responseWriter.WriteHeader(http.StatusInternalServerError)
	model := errorModel{Message: body}
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(model)
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {

}

func AppendOptionedRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), method string) {
	router.HandleFunc(path, handler).Methods(method)
	router.HandleFunc(path, emptyHandler).Methods(http.MethodOptions)
}

func AppendOptionedGetRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request)) {
	AppendOptionedRoute(router, path, handler, http.MethodGet)
}
