package rest

import (
	"net/http"

	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	var pids []string
	for _, v := range g.ListPlayers() {
		pids = append(pids, v.ID())
	}
	logrus.Info("pids:", pids)
	err := ws.PickManager(c.Request.Context()).BroadCast(g.StateJSON(), pids...)
	if err != nil {
		logrus.Errorf("broadcast: %v", err)
	}
	c.String(http.StatusOK, "")
}
