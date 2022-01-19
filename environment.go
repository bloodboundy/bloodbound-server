package main

import (
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

func mixManagers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ctx = game.MixManager(ctx, gameManager.Load().(*game.Manager))
		ctx = player.MixManager(ctx, gameManager.Load().(*player.Manager))
		ctx = ws.MixManager(ctx, wsManager.Load().(*ws.Manager))

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
