package global

import "log"

/**
 * @Author: liu zw
 * @Date: 2021/10/19 17:46
 * @File:
 * @Description: //TODO $
 * @Version:
 */

type OnCommand struct {
	ModuleName  string
	Cmd         string   // 命令
	Alias       []string // 命令别名
	Priority    uint
	CmdFunc     func(sendInfo SendMessage)
	AtMe        bool
	ForMe       bool
	Block       bool
	Permissions []string
}

// 注册插件
func (c OnCommand) Registered() {
	log.Printf("加载模块 %s \n", c.ModuleName)
	CmdSlice = append(CmdSlice, c)
}
