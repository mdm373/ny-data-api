package main

import (
	"flag"
	"log"
	"os"
)

type appParams struct {
	servePort    int
	serveTimeout int
	dbUser       string
	dbHost       string
	dbPassword   string
	swaggerDoc   string
	serveHost    string
}

func parseParams() appParams {
	const required = "[required]"
	flagSet := &flag.FlagSet{}
	flagSet.Init("NY Data API Server", flag.ContinueOnError)
	port := flagSet.Int("port", 8000, "the port server should listen on")
	timeout := flagSet.Int("timeout", 15, "timeout in seconds for read/write")
	host := flagSet.String("host", required, "hostname for postgres backend")
	user := flagSet.String("user", "postgres", "username for postgres backend")
	password := flagSet.String("pass", required, "password for postgres backend")
	serveHost := flagSet.String("serveHost", required, "the visible domain of the server eg http://nydata.xyz.com/")
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
		serveHost:    *serveHost,
	}
}
