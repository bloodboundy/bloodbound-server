package main

import (
	"log"
	"net/http"

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

	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print("ws.read: ", err)
			break
		}

		err = handleWSMsg(ws, mt, msg)
		if err != nil {
			log.Print("ws.handle: ", err)
		}
	}
}

func handleWSMsg(ws *websocket.Conn, mt int, msg []byte) error {
	switch mt {
	case websocket.TextMessage:
	default:
		return nil
	}
	return nil
}
