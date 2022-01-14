// main entry
// handle args, logs and setup HTTP server
package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var (
	PATH_WS = flag.String("ws", "/ws", "websocket api path")
)

func main() {
	flag.Parse()

	log.SetFlags(0)

	mux := http.NewServeMux()
	mux.HandleFunc(*PATH_WS, wsMain)

	log.Fatal(http.ListenAndServe(*addr, mux))
}
