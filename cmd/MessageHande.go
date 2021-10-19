package cmd

import (
	"log"
	"strings"
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
		log.Printf("gocqhttp心跳检测：%s\n", msg["meta_event_type"])
	case "message":
		h := HandleCmd{}
		//senderInfo := msg["sender"].(map[string]interface{})
		rawMessage := msg["raw_message"]
		userID := msg["user_id"]
		groupID := msg["group_id"]
		log.Printf("群聊 %v 接收到来自 %s 的消息：%s \n", msg["group_id"], msg["user_id"], msg["raw_message"])
		//h.BotCmd(groupID, userID, rawMessage)
		var atMe bool
		atMe, rawMessage = CQMessageHandle(rawMessage)
		h.BotCmd(groupID, userID, rawMessage, atMe)
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
