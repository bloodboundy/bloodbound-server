package rest

import (
	"net/http"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/gin-gonic/gin"
)

const (
	PATH_KEY_GID = "game_id"
	PATH_KEY_PID = "player_id"
)

// pickPathGame helper to get the game in the path params, handle 404
func pickPathGame(c *gin.Context) *game.Game {
	return pickGame(c, c.Param(PATH_KEY_GID))
}

func pickGame(c *gin.Context, gid string) *game.Game {
	g, ok := game.PickManager(c.Request.Context()).Load(gid)
	if !ok {
		c.String(http.StatusNotFound, "game not found")
		return nil
	}
	return g
}

// pickPathPlayer helper to get the player in the path params, handle 404
func pickPathPlayer(c *gin.Context) *player.Player {
	return pickPlayer(c, c.Param(PATH_KEY_PID))
}

func pickPlayer(c *gin.Context, pid string) *player.Player {
	p, ok := player.PickManager(c.Request.Context()).Load(pid)
	if !ok {
		c.String(http.StatusNotFound, "player not found")
		return nil
	}
	return p
}

// pickPID returns the player id in the header
func pickPID(c *gin.Context) string {
	return c.GetHeader("Authorization")
}

func isPasswordWrong(c *gin.Context, g *game.Game, pwd string) bool {
	if g.IsPrivate() && g.Password() != c.Param("password") {
		c.String(http.StatusForbidden, "wrong password")
		return true
	}
	return false
}
