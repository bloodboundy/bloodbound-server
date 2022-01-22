package main

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/gin-gonic/gin"
)

var (
	gameManager   atomic.Value
	playerManager atomic.Value
	wsManager     atomic.Value
)

func init() {
	gameManager.Store(game.NewManager())
	playerManager.Store(player.NewManager())
	wsManager.Store(ws.NewManager())
}

// mixManagers mix managers into request context
// middleware used in gin
func mixManagers(c *gin.Context) {
	ctx := c.Request.Context()

	ctx = game.MixManager(ctx, gameManager.Load().(*game.Manager))
	ctx = player.MixManager(ctx, playerManager.Load().(*player.Manager))
	ctx = ws.MixManager(ctx, wsManager.Load().(*ws.Manager))

	c.Request = c.Request.WithContext(ctx)

	c.Next()
}

func extractAuthorization(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	token := strings.Split(auth, "Bearer ")
	if len(token) != 2 {
		c.Error(fmt.Errorf("invalid authorization"))
	} else {
		c.Request.Header.Set("Authorization", token[1])
	}
	c.Next()
}
