package plugins

import (
	"ximan/global"
	"ximan/public/utils"
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
	CmdName:     "ping",
	Cmd:         "ping",
	Alias:       []string{"test"},
	Priority:    1,
	AtMe:        false,
	CmdFunc:     CmdPing,
	Permissions: nil,
	Block:       true,
}
var newPing = global.OnCommand{
	CmdName:     "ping1",
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
	b.Send(sendInfo.MessageType, sendInfo.GroupId, sendInfo.UserId, sendInfo.Message, "send_msg")
}

// newPing指令执行函数
func NewCmdPing(sendInfo global.SendMessage) {
	//time.Sleep(10 * time.Second)
	b := utils.BotSendMessage{AutoEscape: false}
	b.Send(sendInfo.MessageType, sendInfo.GroupId, sendInfo.UserId, "[CQ:face,id=14]pong", "send_msg")
}
