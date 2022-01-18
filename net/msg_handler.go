package net

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type meta struct {
	// meta data of a msg
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Req describes the request from client
type Req struct {
	meta

	// req body, when add a new type req, need to add a field here and a register call in init()
	REGISTER *RegisterReq `json:"REGISTER"`
	LOGIN    *LoginReq    `json:"LOGIN"`
}

type Rsp struct {
	meta

	REGISTER *RegisterRsp `json:"REGISTER"`
	LOGIN    *LoginRsp    `json:"LOGIN"`
}

// init, register handler to handlerMap
func init() {
	registerHandler("REGISTER", handleRegisterReq)
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
