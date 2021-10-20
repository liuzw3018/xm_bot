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
 * @Description: //TODO $
 * @Version:
 */

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
