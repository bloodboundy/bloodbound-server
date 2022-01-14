package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bloodboundy/bloodbound-server/net"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func wsMain(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("ws.upgrade: ", err)
		return
	}
	defer ws.Close()

	ctx := r.Context()
	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print("ws.read: ", err)
			break
		}

		err = handleWSMsg(ctx, ws, mt, msg)
		if err != nil {
			log.Print("ws.handle: ", err)
		}
	}
}

func handleWSMsg(ctx context.Context, ws *websocket.Conn, mt int, msg []byte) error {
	switch mt {
	case websocket.TextMessage:
		return net.HandleTextMessage(ctx, ws, msg)
	default:
		return nil
	}
}
