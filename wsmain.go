package main

import (
	"context"
	"net/http"

	"github.com/bloodboundy/bloodbound-server/net"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

func wsMain(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		handleConnError(w, errors.Wrap(err, "ws.upgrade"))
		return
	}
	defer ws.Close()

	ctx := r.Context()
	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			handleConnError(w, errors.Wrap(err, "ws.read"))
			break
		}

		err = handleWSMsg(ctx, ws, mt, msg)
		if err != nil {
			handleConnError(w, errors.Wrap(err, "ws.handle"))
			// error from handleWSMsg cannot terminal ws conn, thus do not break
		}
	}
}

func handleWSMsg(ctx context.Context, ws *websocket.Conn, mt int, msg []byte) error {
	switch mt {
	case websocket.TextMessage:
		return net.HandleTextMessage(mixServerCtx(ctx), ws, msg)
	default:
		return nil
	}
}

func handleConnError(w http.ResponseWriter, e error) {
	log.Error(e)
	_, _ = w.Write([]byte(e.Error()))
}
