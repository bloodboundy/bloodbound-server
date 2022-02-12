// main entry
// handle args, logs and setup HTTP server
package main

import (
	"flag"

	"github.com/bloodboundy/bloodbound-server/config"
	"github.com/bloodboundy/bloodbound-server/rest"
	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_ "go.uber.org/automaxprocs"
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
	// /games/{game_id}
	route.GET("/games/:"+config.PathGID, rest.GetGamesGID)
	route.PATCH("/games/:"+config.PathGID, rest.PatchGamesGID)
	route.DELETE("/games/:"+config.PathGID, rest.DeleteGamesGID)
	route.GET("/games/:"+config.PathGID+"/players", rest.GetGamesGIDPlayers)
	route.POST("/games/:"+config.PathGID+"/players", rest.PostGamesGIDPlayers)
	route.DELETE("/games/:"+config.PathGID+"/players/:"+config.PathPID, rest.DeleteGamesGIDPlayersPID)

	logrus.Fatal(route.Run(*ADDR))
}
