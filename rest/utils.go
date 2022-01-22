package rest

import (
	"net/http"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/gin-gonic/gin"
)

// pickGame helper to get the game in the path params, handle 404
func pickGame(c *gin.Context) *game.Game {
	g, ok := game.PickManager(c.Request.Context()).Load(c.Param("game_id"))
	if !ok {
		c.String(http.StatusNotFound, "game not found")
		return nil
	}
	return g
}

// pickPlayer helper to get the player in the path params, handle 404
func pickPlayer(c *gin.Context) *player.Player {
	p, ok := player.PickManager(c.Request.Context()).Load(c.Param("player_id"))
	if !ok {
		c.String(http.StatusNotFound, "player not found")
		return nil
	}
	return p
}