package cmd

import (
	"fmt"
	"log"
	"regexp"
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
		atMe, rawMessage := cqMessageHandle(msg["raw_message"], msg["self_id"])
		h.BotCmd(msg["group_id"], msg["user_id"], rawMessage, atMe)
	}
}

func cqMessageHandle(message, selfId interface{}) (bool, string) {
	var atMe bool
	msg := message.(string)
	var returnMsg string
	if strings.Contains(msg, "CQ") {
		temp := strings.Split(msg, " ")
		//log.Println(temp[0])
		if len(temp) == 1 {
			returnMsg = " "
			atMe = cqAtMe(selfId, msg, atMe)
		} else {
			returnMsg = temp[len(temp)-1]
			atMe = cqAtMe(selfId, msg, atMe)
		}
	} else {
		atMe, returnMsg = atMeMessageHandle(msg)
	}
	return atMe, returnMsg
}

func cqAtMe(selfId interface{}, msg string, atMe bool) bool {
	me := fmt.Sprintf("%s", selfId)
	if strings.Contains(msg, "at") && strings.Contains(msg, me) {
		atMe = true
	}
	return atMe
}

func atMeMessageHandle(returnMsg string) (bool, string) {
	var atMe bool
	var reNickName string
	for index, nickName := range global.GConfig.NickName {
		if index == 0 {
			reNickName = fmt.Sprintf("^%s", nickName)
		} else {
			reNickName = fmt.Sprintf("%s|%s", reNickName, nickName)
		}
	}
	matched, _ := regexp.Compile(reNickName)
	if matched.MatchString(returnMsg) {
		atMe = true
		msgLength := len(matched.FindString(returnMsg))
		returnMsg = returnMsg[msgLength:]
	}

	return atMe, returnMsg
}
