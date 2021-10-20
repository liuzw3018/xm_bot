package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"ximan/cmd"
	"ximan/global"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 15:27
 * @File:
 * @Description: 处理CQHttp的消息
 * @Version:
 */

// @title:    	  BotReceiveMessageBase
// @description:  接收CQHttp推送的消息
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func BotReceiveMessageBase(c *gin.Context) {
	var (
		res    []byte
		err    error
		decode *json.Decoder
		msg    = make(map[string]interface{})
	)
	if res, err = ioutil.ReadAll(c.Request.Body); err != nil {
		global.ErrorLogMsgChan <- err.Error()
		return
	}

	decode = json.NewDecoder(strings.NewReader(string(res)))
	decode.UseNumber()

	if err = decode.Decode(&msg); err != nil {
		global.ErrorLogMsgChan <- err.Error()
		return
	}

	cmd.MessageHandle(msg)
}
