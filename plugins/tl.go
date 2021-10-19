package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ximan/global"
	"ximan/public/utils"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/19 15:20
 * @File:
 * @Description: //TODO $
 * @Version:
 */

var tl = global.OnCommand{
	CmdName:     "天行机器人",
	Cmd:         "",
	Alias:       nil,
	AtMe:        true,
	Priority:    10,
	ForMe:       true,
	Permissions: nil,
	CmdFunc:     TlBot,
	Block:       true,
}

type tlBotResp struct {
	Code     int                      `json:"code"`
	Msg      string                   `json:"msg"`
	NewsList []map[string]interface{} `json:"newslist"`
}

func init() {
	tl.Registered()
}

func TlBot(sendInfo global.SendMessage) {
	var (
		question = sendInfo.Message
		userId   = sendInfo.UserId
		key      = "470d505ba2e5dfe67a247ed8cd3e8e9f"
		tlUrl    = fmt.Sprintf("http://api.tianapi.com/txapi/robot/index?key=%s&question=%s&uniqueid=%s&mode=0&priv=0&restype=0", key, question, userId)
		resp     *http.Response
		err      error
		bytes    []byte
		data     tlBotResp
	)

	if resp, err = http.Get(tlUrl); err != nil {
		log.Fatalln(err)
		return
	}
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
		return
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		log.Fatalln(err)
		return
	}

	reply := data.NewsList[0]["reply"]
	b := utils.BotSendMessage{AutoEscape: true}
	b.Send(sendInfo.MessageType, sendInfo.GroupId, sendInfo.UserId, reply, "send_msg")
}
