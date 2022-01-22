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
		c.String(http.StatusBadRequest, err.Error())
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
		if game.IsPrivate() {
			continue
		}
		result = append(result, game.Dump())
	}
	c.JSON(200, result)
}

func GetGamesGID(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	if !isPassed(c, g, c.Param("password")) {
		return
	}
	c.JSON(http.StatusOK, g.Dump())
}

func PatchGamesGID(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	if !isOwner(c, g) {
		return
	}

	var b game.GameJSON
	if c.BindJSON(&b) != nil {
		return
	}
	if err := g.Load(&b); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, g.Dump())
}

func DeleteGamesGID(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	if !isOwner(c, g) {
		return
	}

	game.PickManager(c.Request.Context()).Delete(g.ID())
	c.String(http.StatusOK, "")
}

func GetGamesGIDPlayers(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	if !isPassed(c, g, c.Param("password")) {
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
		ID       string `json:"id"`
		Password string `json:"password"`
	}
	g := pickPathGame(c)
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

	p := pickPlayer(c, rb.ID)
	if p == nil {
		return
	}

	if rb.ID == uid {
		if !isPassed(c, g, rb.Password) {
			return
		}
		if err := g.AddPlayer(p); err != nil {
			c.String(http.StatusInternalServerError, errors.Wrap(err, "AddPlayer").Error())
			return
		}
	} else {
		logrus.Error("TODO: invite")
	}
	c.String(http.StatusOK, "")
}

func DeleteGamesGIDPlayersPID(c *gin.Context) {
	g := pickPathGame(c)
	if g == nil {
		return
	}
	p := pickPathPlayer(c)
	if p == nil {
		return
	}

	if p.ID() == pickPID(c) {
		g.RemovePlayer(p)
	} else {
		logrus.Error("TODO: kick")
	}

	c.String(http.StatusOK, "")
}
