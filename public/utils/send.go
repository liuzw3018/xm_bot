package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

func (b *BotSendMessage) Send(messageType, groupID, userID, message interface{}, msgTag string) {
	switch msgTag {
	case "send_private_msg":
		b.SendPrivateMessage(userID, message)
	case "send_group_msg":
		b.SendGroupMessage(groupID, message)
	case "send_msg":
		b.SendMessage(messageType, groupID, userID, message)
	}
}

func (b *BotSendMessage) SendPrivateMessage(userID, message interface{}) {
	var (
		sendURL = fmt.Sprintf("http://127.0.0.1:5700/send_private_msg?user_id=%s&message=%s", userID, message)
		data    = make(map[string]interface{})
	)
	data = b.postData(sendURL)
	if data["status"] == "ok" {
		log.Println("发送成功:", data)
	} else {
		log.Println(data)
	}

}

func (b *BotSendMessage) SendGroupMessage(groupID, message interface{}) {
	var (
		sendURL = fmt.Sprintf("http://127.0.0.1:5700/send_group_msg?group_id=%s&auto_escape=%s&message=%s", groupID, b.AutoEscape, message)
		data    = make(map[string]interface{})
	)
	data = b.postData(sendURL)
	if data["status"] == "ok" {
		log.Println("发送成功:", data)
	} else {
		log.Println(data)
	}

}

func (b *BotSendMessage) SendMessage(messageType, groupID, userID, message interface{}) {
	var (
		sendURL = fmt.Sprintf("http://127.0.0.1:5700/send_msg?group_id=%s&auto_escape=%v&message=%s&message_type=%s&user_id=%s", groupID, b.AutoEscape, message, messageType, userID)
		data    = make(map[string]interface{})
	)
	data = b.postData(sendURL)
	if data["status"] == "ok" {
		log.Println("发送成功:", data)
	} else {
		log.Println(data)
	}
}

func (b *BotSendMessage) postData(sendURL string) (data map[string]interface{}) {
	var (
		resp  *http.Response
		bytes []byte
		err   error
	)
	resp, err = http.Get(sendURL)
	if err != nil {
		log.Fatalln(err)
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

func (b *BotSendMessage) saveMessageId(messageId interface{}) {

}
