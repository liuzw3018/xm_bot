package plugins

import (
	"log"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/19 20:21
 * @File:
 * @Description: 加载插件
 * @Version:
 */

// @title:    	  LoadPlugin
// @description:  加载插件
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func LoadPlugin() {
	log.Println("模块加载完成！")
	global.BotInfoLogMsgChan <- "模块加载完成！"
}
