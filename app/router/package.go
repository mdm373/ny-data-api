package router

import (
	"github.com/gorilla/mux"
)

//Get The root router for the app
func Get() *mux.Router {
	root := mux.NewRouter()

	appendPrecinct(root)

	return root
}
