package cmd

import (
	"fmt"
	"log"
	"strings"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 16:43
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func MessageHandle(msg map[string]interface{}) {
	switch msg["post_type"] {
	case "meta_event":
		logMsg := fmt.Sprintf("gocqhttp心跳检测：%s", msg["meta_event_type"])
		log.Println(logMsg)
	case "message":
		h := HandleCmd{}
		//senderInfo := msg["sender"].(map[string]interface{})
		logMsg := fmt.Sprintf("群聊 %v 接收到来自 %s 的消息：%s", msg["group_id"], msg["user_id"], msg["raw_message"])
		log.Printf(logMsg)
		global.BotInfoLogMsgChan <- logMsg
		atMe, rawMessage := CQMessageHandle(msg["raw_message"])
		h.BotCmd(msg["group_id"], msg["user_id"], rawMessage, atMe)
	}
}

func CQMessageHandle(message interface{}) (bool, string) {
	msg := message.(string)
	var returnMsg string
	if strings.Contains(msg, "CQ") {
		temp := strings.Split(msg, " ")
		//log.Println(temp[0])
		if len(temp) == 1 {
			returnMsg = " "
		} else {
			returnMsg = temp[len(temp)-1]
		}
	} else {
		returnMsg = msg
	}

	return strings.Contains(msg, "at") && strings.Contains(msg, "535778382"), returnMsg
}
