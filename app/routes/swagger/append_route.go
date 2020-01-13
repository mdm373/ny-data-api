package swagger

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AppendRoute(parent *mux.Router, host string) {
	fs := http.FileServer(http.Dir("./static/swagger-ui/"))
	parent.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", fs))
	parent.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fmt.Sprintf("%s/swagger-ui/", host), http.StatusFound)
	})
}
