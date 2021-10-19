package server

import (
	"github.com/gin-gonic/gin"
	v1 "ximan/api/v1"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/13 9:11
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func RunServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())

	app.POST("/gocqhttp", v1.BotReceiveMessageBase)

	return app
}
