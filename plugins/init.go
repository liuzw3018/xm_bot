package plugins

import (
	"log"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/19 20:21
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func LoadPlugin() {
	log.Println("模块加载完成！")
	global.BotInfoLogMsgChan <- "模块加载完成！"
}
