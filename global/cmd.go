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

// OnCommand 事件结构体
type OnCommand struct {
	ModuleName  string                     // 插件名称
	Cmd         string                     // 事件名称
	Aliases     []string                   // 事件别名
	Priority    uint                       // 优先级
	CmdFunc     func(sendInfo SendMessage) // 事件执行函数
	AtMe        bool                       // 是否at机器人
	ForMe       bool                       // 是否回应所有的消息
	Block       bool                       // 是否继续匹配下一个事件
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
