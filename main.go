package main

import (
	"fmt"
	"ximan/global"
	"ximan/server"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/11 14:50
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func main() {
	app := server.RunServer()

	addr := fmt.Sprintf("%s:%d", global.GConfig.Host, global.GConfig.Port)
	err := app.Run(addr)
	if err != nil {
		panic(err)
	}
}
