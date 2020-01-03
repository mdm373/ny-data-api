package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mdm373/ny-data-api/app/bounds"
	"github.com/mdm373/ny-data-api/app/series"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdm373/ny-data-api/app/db"
)

type appParams struct {
	servePort    int
	serveTimeout int
	dbUser       string
	dbHost       string
	dbPassword   string
}

func main() {
	params := parseParams()
	connection := connect(params.dbUser, params.dbPassword, params.dbHost)
	serve(params.servePort, params.serveTimeout, connection)
	log.Print("app ended")
}

const required = "[required]"

func parseParams() appParams {
	flagSet := &flag.FlagSet{}
	flagSet.Init("NY Data API Server", flag.ContinueOnError)
	port := flagSet.Int("port", 8000, "the port server should listen on")
	timeout := flagSet.Int("timeout", 15, "timeout in seconds for read/write")
	host := flagSet.String("host", required, "hostname for postgres backend")
	user := flagSet.String("user", "postgres", "username for postgres backend")
	password := flagSet.String("pass", required, "password for postgres backend")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		log.Printf("command line exit - %s", err)
		os.Exit(0)
	}
	if *user == required || *host == required {
		log.Printf("host and pass are required params")
		os.Exit(0)
	}
	return appParams{
		servePort:    *port,
		serveTimeout: *timeout,
		dbHost:       *host,
		dbPassword:   *password,
		dbUser:       *user,
	}
}

func serve(port int, timeout int, connection db.Connection) {
	log.Printf("serving on port %d\n", port)
	handler, err := getRouter(connection)
	if err != nil {
		if err != nil {
			log.Printf("server start fail - %+v", err)
			return
		}
	}
	server := &http.Server{
		Handler:      handler,
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("server exit - %+v", err)
		return
	}
}

func getRouter(connection db.Connection) (*mux.Router, error) {
	root := mux.NewRouter().StrictSlash(true)
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
