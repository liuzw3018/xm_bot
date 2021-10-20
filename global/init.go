package global

import (
	"log"
	"ximan/config"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/20 9:23
 * @File:
 * @Description: //TODO $
 * @Version:
 */

var (
	GConfig = config.Config{}
)

func init() {
	log.Println("加载配置文件...")
	GConfig = config.LoadConfig()
}
