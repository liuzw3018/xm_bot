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
 * @Description: 对从api接收到的消息进行预处理
 * @Version:
 */

// @title:    	  MessageHandle
// @description:  对从api接收到的消息进行预处理
// @auth:         liuzw3018
// @param:        msg map[string]interface{}
// @return:       nil
func MessageHandle(msg map[string]interface{}) {
	// TODO 匹配消息类型
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

// @title:    	  cqMessageHandle
// @description:  处理消息中包含的CQ码信息，并判断是否at了机器人
// @auth:         liuzw3018
// @param:        message, selfId interface{}
// @return:       bool, string
func cqMessageHandle(message, selfId interface{}) (bool, string) {
	var atMe bool
	msg := message.(string)
	var returnMsg string
	// TODO 如果消息中含有 CQ字符，需要分割消息，否则去匹配 nickname
	if strings.Contains(msg, "CQ") {
		temp := strings.Split(msg, " ")
		//log.Println(temp[0])
		// TODO 分割消息之后切片长度为 1，消息中只含有 CQ码消息，判断消息是否为 at机器人
		if len(temp) == 1 {
			returnMsg = " "
			atMe = cqAtMe(selfId, msg, atMe)
		} else {
			returnMsg = temp[len(temp)-1]
			atMe = cqAtMe(selfId, msg, atMe)
		}
	} else {
		// TODO 匹配nickname
		atMe, returnMsg = atMeMessageHandle(msg)
	}
	return atMe, returnMsg
}

// @title:    	  cqAtMe
// @description:  判断消息中包含的CQ码是否为at机器人
// @auth:         liuzw3018
// @param:        selfId interface{}, msg string, atMe bool
// @return:       bool
func cqAtMe(selfId interface{}, msg string, atMe bool) bool {
	me := fmt.Sprintf("%s", selfId)
	// TODO CQ码中含有 at字符并且含有自己的 QQ号，返回 true
	if strings.Contains(msg, "at") && strings.Contains(msg, me) {
		atMe = true
	}
	return atMe
}

// @title:    	  atMeMessageHandle
// @description:  响应消息中机器人的nickname，判定为at机器人
// @auth:         liuzw3018
// @param:        returnMsg string
// @return:       bool, string
func atMeMessageHandle(returnMsg string) (bool, string) {
	var atMe bool
	var reNickName string
	// TODO nickname 添加到正则表达式中
	for index, nickName := range global.GConfig.NickName {
		if index == 0 {
			reNickName = fmt.Sprintf("^%s", nickName)
		} else {
			reNickName = fmt.Sprintf("%s|%s", reNickName, nickName)
		}
	}
	matched, _ := regexp.Compile(reNickName)
	// TODO 消息中含有 nickname，截取消息并返回
	if matched.MatchString(returnMsg) {
		atMe = true
		msgLength := len(matched.FindString(returnMsg))
		returnMsg = returnMsg[msgLength:]
	}

	return atMe, returnMsg
}
