package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdm373/ny-data-api/app/db"
	"github.com/mdm373/ny-data-api/app/router"
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
	connect(params.dbUser, params.dbPassword, params.dbHost)
	serve(params.servePort, params.serveTimeout)
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
		log.Printf("user and host are required params")
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

func serve(port int, timeout int) {
	log.Printf("serving on port %d\n", port)
	server := &http.Server{
		Handler:      router.Get(),
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("server exit - %s", err)
		return
	}
}

func connect(user string, password string, host string) {
	connection, err := db.Get(user, password, host)
	if err != nil {
		log.Printf("db connection failure - %s", err)
		return
	}
	log.Printf("connected to db: %s", connection)
}
