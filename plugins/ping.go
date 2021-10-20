package plugins

import (
	"ximan/global"
	"ximan/utils"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 15:53
 * @File:
 * @Description: //TODO $
 * @Version:
 */

// 创建指令实例
var ping = global.OnCommand{
	ModuleName:  "ping",
	Cmd:         "ping",
	Alias:       []string{"test"},
	Priority:    1,
	AtMe:        false,
	CmdFunc:     CmdPing,
	Permissions: nil,
	Block:       true,
}
var newPing = global.OnCommand{
	ModuleName:  "ping1",
	Cmd:         "ping1",
	Alias:       []string{"test1"},
	Priority:    1,
	AtMe:        true,
	CmdFunc:     NewCmdPing,
	Permissions: nil,
	Block:       true,
}

// 注册模块
func init() {
	ping.Registered()
	newPing.Registered()
}

// ping指令执行函数
func CmdPing(sendInfo global.SendMessage) {
	// 实例化消息发送对象
	b := utils.BotSendMessage{AutoEscape: true}
	// 消息发送
	b.Send(sendInfo, "send_msg")
}

// newPing指令执行函数
func NewCmdPing(sendInfo global.SendMessage) {
	//time.Sleep(10 * time.Second)
	sendInfo.Message = "[CQ:face,id=14]pong"
	b := utils.BotSendMessage{AutoEscape: false}
	sendInfo.AtSender = true
	b.Send(sendInfo, "send_msg")
}
