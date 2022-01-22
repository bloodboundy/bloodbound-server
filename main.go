// main entry
// handle args, logs and setup HTTP server
package main

import (
	"flag"

	"github.com/bloodboundy/bloodbound-server/rest"
	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Net flags
var ADDR = flag.String("addr", "localhost:8080", "http service address")

// Log flags
var (
	LEVEL = flag.String("log-level", "info",
		`log level, accepts: trace debug info warning error fatal panic, default: info`)
	OUTPUT = flag.String("log-output", "", `the file log will be output to, default to stderr`)
)

func main() {
	flag.Parse()
	setupLogger()

	route := gin.Default()
	route.Use(mixManagers)
	route.Use(extractAuthorization)

	// /ws
	route.Any("/ws", ws.Main)

	// /register
	route.GET("/register", rest.GetRegister)

	// /games
	route.GET("/games", rest.GetGames)
	route.POST("/games", rest.PostGames)

	logrus.Fatal(route.Run(*ADDR))
}
