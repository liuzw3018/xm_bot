package global

import (
	"fmt"
	"log"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/19 17:46
 * @File:
 * @Description: 生成插件
 * @Version:
 */

// OnCommand 命令结构体，实现了用户插件命令
type OnCommand struct {
	ModuleName  string                     // 插件名称
	Cmd         string                     // 命令
	Aliases     []string                   // 命令别名
	Priority    uint                       // 优先级
	CmdFunc     func(sendInfo SendMessage) // 命令执行函数
	AtMe        bool                       // 是否at机器人
	ForMe       bool                       // 是否回应所有的消息
	Block       bool                       // 是否继续执行下一个命令
	Permissions []string                   // 超级用户
}

// @title:    	  Registered
// @description:  注册插件
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func (c OnCommand) Registered() {
	logMsg := fmt.Sprintf("加载模块 %s", c.ModuleName)
	log.Println(logMsg)
	BotInfoLogMsgChan <- logMsg
	CmdSlice = append(CmdSlice, c)
}
