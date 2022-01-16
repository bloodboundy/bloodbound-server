// main entry
// handle args, logs and setup HTTP server
package main

import (
	"flag"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// Net flags
var ADDR = flag.String("addr", "localhost:8080", "http service address")

var (
	PATH_WS = flag.String("ws", "/ws", "websocket api path")
)

// Log flags
var (
	LEVEL = flag.String("log-level", "info",
		`log level, accepts: trace debug info warning error fatal panic, default: info`)
	OUTPUT = flag.String("log-output", "", `the file log will be output to, default to stderr`)
)

func main() {
	flag.Parse()

	setupLogger()

	mux := http.NewServeMux()
	mux.HandleFunc(*PATH_WS, wsMain)

	log.Fatal(http.ListenAndServe(*ADDR, mux))
}

func setupLogger() {
	var level log.Level
	err := level.UnmarshalText([]byte(*LEVEL))
	if err != nil {
		log.Panicf("unaccptable level: %v, accepts: trace debug info warning error fatal panic", *LEVEL)
	}
	log.SetLevel(level)

	out := os.Stderr
	if *OUTPUT != "" {
		file, err := os.OpenFile(*OUTPUT, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Panicf("failed to open log output file: %v", err)
		}
		out = file
	}
	log.SetOutput(out)
}
