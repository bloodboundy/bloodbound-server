package net

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/gorilla/websocket"
)

func init() {
	registerHandler("REGISTER", handleRegisterReq)
}

type RegisterReq struct {
}

type RegisterRsp player.Player

func handleRegisterReq(ctx context.Context, ws *websocket.Conn, req *Req) error {

	return nil
}
