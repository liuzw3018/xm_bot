package xmevent

import (
	"sort"
	"ximan/global"
	"ximan/plugins"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 16:06
 * @File:
 * @Description: 对优化过的用户信息进行处理
 * @Version:
 */

// @title:    	  init
// @description:  初始化插件，按优先级对插件进行排序
// @auth:         liuzw3018
// @param:        nil
// @return:       nil
func init() {
	plugins.LoadPlugin()
	sort.Sort(global.CmdSlice)
}

// HandleCmd 命令处理对象
type HandleCmd struct {
}

// @title:    	  BotCmd
// @description:  对比插件中的命令，执行用户事件
// @auth:         liuzw3018
// @param:        groupID, userID, rawMessage interface{}, atMe bool
// @return:       nil
func (h *HandleCmd) BotCmd(groupID, userID, rawMessage interface{}, atMe bool) {
	// TODO 遍历插件列表匹配事件
	for _, oc := range global.CmdSlice {
		switch {
		//case rawMessage == "notice":
		case rawMessage == oc.Cmd:
			h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
			if oc.Block {
				return
			}
		case rawMessage != "":
			// TODO 遍历插件列表匹配事件别名
			for _, aliasCmd := range oc.Aliases {
				if rawMessage == aliasCmd {
					h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
					if oc.Block {
						return
					}
				}
			}
			// TODO ForMe类事件如果Block为true，则退出下面的匹配
			if ok := h.forMe(groupID, userID, rawMessage, atMe, oc); ok {
				return
			}
		case rawMessage == "":
			if ok := h.forMe(groupID, userID, rawMessage, atMe, oc); ok {
				return
			}
		}

	}
}

// @title:    	  forMe
// @description:  处理无法匹配的事件
// @auth:         liuzw3018
// @param:        groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand
// @return:       bool
func (h *HandleCmd) forMe(groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand) bool {
	// TODO 如果事件可以接收所有信息则执行
	if oc.ForMe {
		h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
		if oc.Block {
			return true
		}
	}
	return false
}

// @title:    	  atBotMessageHandle
// @description:  对已经匹配到的事件判断是否需要唤醒机器人
// @auth:         liuzw3018
// @param:        groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand
// @return:       nil
func (h *HandleCmd) atBotMessageHandle(groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand) {
	//fmt.Println(oc)
	// TODO 匹配事件响应方式
	switch {
	case atMe: // at机器人，无论如何都要执行
		h.runCommand(groupID, userID, rawMessage, oc)
	case !oc.AtMe: // 不需要at机器人就可以执行
		h.runCommand(groupID, userID, rawMessage, oc)
	case oc.AtMe && atMe: // 需要at机器人并且已经at了机器人才可以执行
		h.runCommand(groupID, userID, rawMessage, oc)
	default:
		return
	}
}

// @title:    	  runCommand
// @description:  对已经完成处理的事件执行对应的函数
// @auth:         liuzw3018
// @param:        groupID, userID, rawMessage interface{}, oc global.OnCommand
// @return:       nil
func (h *HandleCmd) runCommand(groupID, userID, rawMessage interface{}, oc global.OnCommand) {
	sendInfo := global.SendMessage{
		UserId:      userID,
		Message:     rawMessage.(string),
		MessageType: "private",
	}
	// TODO 如果含有群聊信息，则添加群聊相关字段
	if groupID != nil {
		sendInfo.IsGroup = true
		sendInfo.GroupId = groupID
		sendInfo.MessageType = "group"
	}
	oc.CmdFunc(sendInfo)
}
