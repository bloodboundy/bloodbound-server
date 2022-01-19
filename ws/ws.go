package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{}

func Main(c *gin.Context) {
	uid := c.Request.Header.Get("Authorization")
	wm := PickManager(c.Request.Context())
	if _, ok := wm.Load(uid); ok {
		handleConnError(c.Writer, fmt.Errorf("this player is already in ws conn"))
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		handleConnError(c.Writer, errors.Wrap(err, "ws.upgrade"))
		return
	}
	defer ws.Close()

	if err := wm.Store(uid, ws); err != nil {
		handleConnError(c.Writer, errors.Wrap(err, "wm.Store"))
		return
	}

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			wm.Delete(uid)
			handleConnError(c.Writer, errors.Wrap(err, "ws.read"))
			break
		}
	}
}

func handleConnError(w http.ResponseWriter, e error) {
	log.Error(e)
	_, _ = w.Write([]byte(e.Error()))
}
