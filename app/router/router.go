package router

import (
	"encoding/json"
	"net/http"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

func RespondWithJSON(responseWriter http.ResponseWriter, body interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(body)
}
