package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/routes/bounds"
	"github.com/mdm373/ny-data-api/app/routes/series"
	"github.com/mdm373/ny-data-api/app/routes/swagger"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdm373/ny-data-api/app/db"
)

func main() {
	params := parseParams()
	connection := connect(params.dbUser, params.dbPassword, params.dbHost)
	serve(params.servePort, params.serveTimeout, params.serveHost, connection)
	log.Println("app ended")
}

func serve(port int, timeout int, serveHost string, connection db.Connection) {
	log.Printf("serving on port %d\n", port)
	handler, err := getRouter(connection, serveHost)

	if err != nil {
		log.Printf("server start fail - %+v", err)
		return
	}
	server := &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Println("server exit", err)
		return
	}
}

func getRouter(connection db.Connection, serveHost string) (*mux.Router, error) {
	root := mux.NewRouter().StrictSlash(true)
	swagger.AppendRoute(root, serveHost)
	err := bounds.AppendRoute(root, connection)
	if err != nil {
		return nil, err
	}
	err = series.AppendRoute(root, connection)
	if err != nil {
		return nil, err
	}
	root.Use(mux.CORSMethodMiddleware(root))
	root.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})
	return root, nil
}

func connect(user string, password string, host string) db.Connection {
	connection, err := db.Get(user, password, host)
	if err != nil {
		log.Printf("db connection failure - %+v", err)
		os.Exit(0)
	}
	log.Printf("connected to db: %+v", connection)
	return connection
}
