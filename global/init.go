package global

import (
	"log"
	"ximan/config"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/20 9:23
 * @File:
 * @Description: 初始化项目所需的变量
 * @Version:
 */

var (
	GConfig = &config.Config{} // 全局配置对象
)

// @title:    	  init
// @description:  初始化项目所需的变量
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func init() {
	log.Println("加载配置文件...")
	BotInfoLogMsgChan <- "加载配置文件..."
	GConfig = config.LoadConfig()
}
