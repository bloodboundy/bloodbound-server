package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PatchGamesGIDStatus(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	if err := g.Start(); err != nil {
		c.String(http.StatusInternalServerError, "game start failed: %v", err)
		return
	}
	c.String(http.StatusOK, "")
}
