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
 * @Description: //TODO $
 * @Version:
 */

//var messageIdSlice []interface{}

type BotSendMessage struct {
	AutoEscape bool // 消息内容是否作为纯文本发送 ( 即不解析 CQ 码 ) , 只在 message 字段是字符串时有效
}

func (b *BotSendMessage) Send(sendInfo global.SendMessage, msgTag string) {
	switch msgTag {
	case "send_private_msg":
		b.SendPrivateMessage(sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	case "send_group_msg":
		b.SendGroupMessage(sendInfo.GroupId, sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	case "send_msg":
		b.SendMessage(sendInfo.MessageType, sendInfo.GroupId, sendInfo.UserId, sendInfo.Message, sendInfo.AtSender)
	}
}

func (b *BotSendMessage) SendPrivateMessage(userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_private_msg?user_id=%s&message=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, userId, newMsg)
	data = b.postData(sendURL)
	b.checkSendStatus(data)

}

func (b *BotSendMessage) SendGroupMessage(groupId, userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_group_msg?group_id=%s&auto_escape=%v&message=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, groupId, b.AutoEscape, newMsg)
	data = b.postData(sendURL)
	b.checkSendStatus(data)
}

func (b *BotSendMessage) SendMessage(messageType, groupId, userId, message interface{}, atSender bool) {
	var (
		sendURL string
		data    = make(map[string]interface{})
		newMsg  interface{}
	)
	if atSender {
		newMsg = atMessageHandle(message, userId)
	} else {
		newMsg = message
	}
	sendURL = fmt.Sprintf("http://%s:%d/send_msg?group_id=%s&auto_escape=%v&message=%s&message_type=%s&user_id=%s", global.GConfig.CQHttp.Host, global.GConfig.CQHttp.Port, groupId, b.AutoEscape, newMsg, messageType, userId)
	data = b.postData(sendURL)
	b.checkSendStatus(data)
}

func (b *BotSendMessage) postData(sendURL string) (data map[string]interface{}) {
	var (
		resp  *http.Response
		bytes []byte
		err   error
	)
	resp, err = http.Get(sendURL)
	if err != nil {
		log.Println(err)
		return nil
	}
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println(err)
		return nil
	}
	decode := json.NewDecoder(strings.NewReader(string(bytes)))
	decode.UseNumber()
	err = decode.Decode(&data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return data
}

func (b *BotSendMessage) checkSendStatus(data map[string]interface{}) {
	if data["status"] == "ok" {
		log.Println("发送成功:", data)
	} else {
		log.Println(data)
	}
}

func (b *BotSendMessage) saveMessageId(messageId interface{}) {

}

func atMessageHandle(message, userId interface{}) string {
	return fmt.Sprintf("[CQ:at,qq=%s]%s", userId, message)
}
