package server

import (
	"github.com/gin-gonic/gin"
	v1 "ximan/api/v1"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/13 9:11
 * @File:
 * @Description: 定义机器人服务器
 * @Version:
 */

// @title:    	  RunServer
// @description:  实例化gin服务器引擎，定义url
// @auth:         liuzw3018
// @param:        nil
// @return:       *gin.Engine
func RunServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())

	app.POST("/gocqhttp", v1.BotReceiveEventBase)

	return app
}
