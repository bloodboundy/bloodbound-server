package rest

import (
	"github.com/bloodboundy/bloodbound-server/player"
	"github.com/gin-gonic/gin"
)

func GetRegister(c *gin.Context) {
	nickname := c.Query("nickname")
	if nickname == "" {
		c.String(400, "nickname required, used as /register?nickname=xxx")
		return
	}

	p, err := player.PickManager(c.Request.Context()).Register(nickname)
	if err != nil {
		c.String(500, "Register: %v", err)
		return
	}

	c.JSON(200, p)
}
