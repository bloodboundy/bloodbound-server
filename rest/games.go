package rest

import (
	"net/http"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func PostGames(c *gin.Context) {
	var b game.GameJSON
	if c.BindJSON(&b) != nil {
		return
	}

	g := game.NewGame(pickPID(c))
	if err := g.Load(&b); err != nil {
		c.String(500, "Load: %v", err)
		return
	}
	_, loaded := game.PickManager(c.Request.Context()).LoadOrStore(g)
	if loaded {
		c.String(500, "game id dup")
		return
	}
	c.JSON(200, g.Dump("password"))
}

func GetGames(c *gin.Context) {
	games := game.PickManager(c.Request.Context()).List()
	result := []*game.GameJSON{}
	for _, game := range games {
		result = append(result, game.Dump())
	}
	c.JSON(200, result)
}

func GetGamesGID(c *gin.Context) {
	g := pickGame(c)
	if g == nil {
		return
	}
	c.JSON(200, g.Dump())
}

func GetGamesGIDPlayers(c *gin.Context) {
	g := pickGame(c)
	if g == nil {
		return
	}
	players := []*player.PlayerJSON{}
	for _, p := range g.ListPlayers() {
		players = append(players, p.Dump())
	}
	c.JSON(200, players)
}

func PostGamesGIDPlayers(c *gin.Context) {
	type reqBody struct {
		ID string `json:"id"`
	}
	g := pickGame(c)
	if g == nil {
		return
	}

	var rb reqBody
	if c.BindJSON(&rb) != nil {
		return
	}

	uid := pickPID(c)
	if rb.ID == "" {
		rb.ID = uid
	}

	p, ok := player.PickManager(c.Request.Context()).Load(rb.ID)
	if !ok {
		c.String(http.StatusNotFound, "player not found")
		return
	}

	if rb.ID == uid {
		if err := g.AddPlayer(p); err != nil {
			c.String(http.StatusInternalServerError, errors.Wrap(err, "AddPlayer").Error())
			return
		}
	} else {
		logrus.Error("TODO: invite")
	}
	c.String(http.StatusOK, "")
}
