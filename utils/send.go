package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 15:31
 * @File:
 * @Description: 消息发送
 * @Version:
 */

//var messageIdSlice []interface{}
// BotSendMessage 机器人消息发送方法
type BotSendMessage struct {
	AutoEscape bool // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
}

// @title:    	  Send
// @description:  发送消息，并根据tag判断发送消息的类型
// @auth:         liuzw3018
// @param:        sendInfo global.SendMessage, msgTag string
// @return:       nil
func (b *BotSendMessage) Send(sendInfo global.SendMessage, msgTag string) {
	// TODO 匹配消息 tag
	switch msgTag {
	case "send_private_msg":
		b.SendPrivateMessage(sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	case "send_group_msg":
		b.SendGroupMessage(sendInfo.GroupId, sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	case "send_msg":
		b.SendMessage(sendInfo.MessageType, sendInfo.GroupId, sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	}
}

// @title:    	  SendPrivateMessage
// @description:  发送私聊消息
// @auth:         liuzw3018
// @param:        userId, message interface{}, atSender bool
// @return:       nil
func (b *BotSendMessage) SendPrivateMessage(userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	// TODO 判断是否需要 at发送人
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_private_msg?user_id=%s&message=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, userId, newMsg)
	data = b.postData(sendURL)
	b.checkSendStatus(data)

}

// @title:    	  SendGroupMessage
// @description:  发送群聊消息
// @auth:         liuzw3018
// @param:        userId, message interface{}, atSender bool
// @return:       nil
func (b *BotSendMessage) SendGroupMessage(groupId, userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	// TODO 判断是否需要 at发送人
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_group_msg?group_id=%s&auto_escape=%v&message=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, groupId, b.AutoEscape, newMsg)
	data = b.postData(sendURL)
	b.checkSendStatus(data)
}

// @title:    	  SendMessage
// @description:  发送消息，messageType判断发送私聊消息还是群聊消息
// @auth:         liuzw3018
// @param:        messageType, groupId, userId, message interface{}, atSender bool
// @return:       nil
func (b *BotSendMessage) SendMessage(messageType, groupId, userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	// TODO 判断是否需要 at发送人
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_msg?group_id=%s&auto_escape=%v&message=%s&message_type=%s&user_id=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, groupId, b.AutoEscape, newMsg, messageType, userId)
	data = b.postData(sendURL)
	b.checkSendStatus(data)
}

// @title:    	  postData
// @description:  推送消息到CQHttp发送
// @auth:         liuzw3018
// @param:        sendURL string
// @return:       data map[string]interface{}
func (b *BotSendMessage) postData(sendURL string) (data map[string]interface{}) {
	var (
		resp  *http.Response
		bytes []byte
		err   error
	)
	resp, err = http.Get(sendURL)
	if err != nil {
		err = fmt.Errorf("get请求出错：%s", err)
		log.Println(err)
		global.ErrorLogMsgChan <- err.Error()
		return nil
	}
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("读取body出错：%s", err)
		log.Println(err)
		global.ErrorLogMsgChan <- err.Error()
		return nil
	}
	decode := json.NewDecoder(strings.NewReader(string(bytes)))
	decode.UseNumber()
	err = decode.Decode(&data)
	if err != nil {
		err = fmt.Errorf("json反序列化出错：%s", err)
		log.Println(err)
		global.ErrorLogMsgChan <- err.Error()
		return nil
	}
	return data
}

// @title:    	  checkSendStatus
// @description:  检查消息发送状态
// @auth:         liuzw3018
// @param:        data map[string]interface{}
// @return:       nil
func (b *BotSendMessage) checkSendStatus(data map[string]interface{}) {
	if data["status"] == "ok" {
		log.Println("发送成功:", data)
	} else {
		log.Println("消息发送失败：", data)
	}
}

// @title:    	  saveMessageId
// @description:  保存上一次发送的消息ID(未完成)
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func (b *BotSendMessage) saveMessageId(messageId interface{}) {

}

// @title:    	  atMessageHandle
// @description:  添加at发送人消息头
// @auth:         liuzw3018
// @param:        message, userId interface{}
// @return:       string
func atMessageHandle(message, userId interface{}) string {
	return fmt.Sprintf("[CQ:at,qq=%s]%s", userId, message)
}
