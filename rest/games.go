package rest

import (
	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func PostGames(c *gin.Context) {
	var b game.GameJSON
	if err := c.ShouldBindJSON(&b); err != nil {
		c.String(500, errors.Wrap(err, "BindJSON").Error())
	}

	uid := c.GetHeader("Authorization")
	gm := game.PickManager(c.Request.Context())
	g := gm.NewGame(uid)
	if err := g.Load(&b); err != nil {
		gm.Delete(g.ID)
		c.String(500, "Load: %v", err)
		return
	}

	c.JSON(200, g.Dump("players", "password"))
}

func GetGames(c *gin.Context) {
	games := game.PickManager(c.Request.Context()).List()
	result := []*game.GameJSON{}
	for _, game := range games {
		result = append(result, game.Dump("players"))
	}
	c.JSON(200, result)
}
