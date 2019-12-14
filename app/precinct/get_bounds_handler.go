package precinct

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
)

const idPathParam = "id"

type exampleResponse struct {
	Message string `json:"message"`
}

func newGetBoundsHandler(connection db.Connection) router.RouteHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if _, ok := vars[idPathParam]; ok {
			router.RespondWithJSON(w, exampleResponse{Message: "hello world"})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

}
