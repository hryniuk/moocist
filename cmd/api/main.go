package main

import (
	"log"
	"net/http"

	"github.com/hryniuk/moocist/pkg/api"
	"github.com/sirupsen/logrus"
)

const addr = "127.0.0.1:8181"

func initLogging() {
	logrus.SetLevel(logrus.DebugLevel)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
}

func main() {
	initLogging()
	server := api.NewServer()

	httpServer := &http.Server{
		Addr:    addr,
		Handler: server.Router,
	}

	log.Printf("starting server at %s", addr)
	log.Fatal(httpServer.ListenAndServe())
}
