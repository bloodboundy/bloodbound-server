package net

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Req struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	REGISTER *RegisterReq `json:"REGISTER"`
	LOGIN    *LoginReq    `json:"LOGIN"`
}

type handler func(ctx context.Context, ws *websocket.Conn, req *Req) error

var handlerMap = map[string]handler{}

func registerHandler(k string, h handler) {
	_, ok := handlerMap[k]
	if ok {
		panic(fmt.Sprintf("handler for %s is already registered", k))
	}
	handlerMap[k] = handleRegisterReq
}

func HandleTextMessage(ctx context.Context, ws *websocket.Conn, msg []byte) error {
	req := &Req{}

	err := json.Unmarshal(msg, req)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}

	h, ok := handlerMap[req.Type]
	if !ok {
		return fmt.Errorf("unknown req type: %v", req.Type)
	}
	err = h(ctx, ws, req)
	if err != nil {
		return errors.Wrap(err, "handle")
	}
	return nil
}
