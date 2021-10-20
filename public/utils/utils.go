package utils

import (
	"log"
	"ximan/global"
	"ximan/xmlog"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 16:17
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func ErrorLogMsg() {
	for {
		if v, ok := <-global.ErrorLogMsgChan; ok {
			log.Println(v)
			err := xmlog.WriteLog("xmBotError.log", v)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func XmInfoLog() {
	for {
		if v, ok := <-global.BotInfoLogMsgChan; ok {
			err := xmlog.WriteLog("xmBotInfo.log", v)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
