package main

import (
	"fmt"
	"ximan/global"
	"ximan/server"
	"ximan/utils"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/11 14:50
 * @File:
 * @Description: 主函数
 * @Version:
 */

// @title:    	  main
// @description:  主函数，机器人从这里启动
// @auth:         liuzw3018
// @param:        fileName, msg string
// @return:       error
func main() {
	app := server.RunServer()
	// 启动日志处理
	go utils.ErrorLogMsg()
	go utils.XmInfoLog()
	// 启动服务器，监听端口
	addr := fmt.Sprintf("%s:%d", global.GConfig.Host, global.GConfig.Port)
	err := app.Run(addr)
	if err != nil {
		panic(err)
	}
}
