package main

import "ximan/server"

/**
 * @Author: liu zw
 * @Date: 2021/10/11 14:50
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func main() {
	app := server.RunServer()

	err := app.Run(":5701")
	if err != nil {
		panic(err)
	}
}
