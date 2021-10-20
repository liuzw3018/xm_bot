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
	ModuleName:  "天行机器人",
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
		tlUrl    = fmt.Sprintf("%s?key=%s&question=%s&uniqueid=%s&mode=0&priv=0&restype=0", global.GConfig.TxBotUrl, global.GConfig.TxBotKey, question, userId)
		resp     *http.Response
		err      error
		bytes    []byte
		data     tlBotResp
		b        = utils.BotSendMessage{AutoEscape: true}
		reply    interface{}
	)
	if resp, err = http.Get(tlUrl); err != nil {
		log.Println("get请求出错：", err)
		sendInfo.Message = err
		goto END
	}
	if bytes, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Println("读取body错误：", err)
		sendInfo.Message = err
		goto END
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		log.Println("json反序列化错误：", err)
		sendInfo.Message = err
		goto END
	}

	reply = data.NewsList[0]["reply"]
	sendInfo.Message = reply
END:
	b.Send(sendInfo, "send_msg")
}
