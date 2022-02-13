package rest

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostGamesGIDActions(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	req := c.Request.Body
	defer req.Close()
	data, err := io.ReadAll(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "read reqbody failed")
		return
	}
	if err := g.ApplyAction(c.Request.Context(), data); err != nil {
		c.String(http.StatusBadRequest, "apply action: %v", err)
	}
	c.String(http.StatusOK, "")
}
