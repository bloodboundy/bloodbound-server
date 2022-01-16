package main

import (
	"context"
	"sync/atomic"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/player"
)

var (
	gameManager   atomic.Value
	playerManager atomic.Value
)

func init() {
	gameManager.Store(game.NewManager())
	playerManager.Store(player.NewManager())

}

func mixServerCtx(ctx context.Context) context.Context {
	ctx = game.MixManager(ctx, gameManager.Load().(*game.Manager))
	ctx = player.MixManager(ctx, playerManager.Load().(*player.Manager))
	return ctx
}
