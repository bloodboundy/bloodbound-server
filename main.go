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

	log.Fatal(route.Run(*ADDR))
}
