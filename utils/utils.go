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
 * @Description: 日志处理
 * @Version:
 */

// @title:    	  ErrorLogMsg
// @description:  处理错误日志
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
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

// @title:    	  XmInfoLog
// @description:  处理普通日志
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
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
