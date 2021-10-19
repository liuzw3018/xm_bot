package cmd

import (
	"sort"
	"ximan/global"
	"ximan/plugins"
)

/**
 * @Author: liu zw
 * @Date: 2021/10/18 16:06
 * @File:
 * @Description: //TODO $
 * @Version:
 */

func init() {
	plugins.LoadPlugin()
	sort.Sort(global.CmdSlice)
}

type HandleCmd struct {
}

func (h *HandleCmd) BotCmd(groupID, userID, rawMessage interface{}, atMe bool) {
	for _, oc := range global.CmdSlice {
		switch {
		case rawMessage == oc.Cmd:
			h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
			if oc.Block {
				return
			}
		case rawMessage != "":
			for _, aliasCmd := range oc.Alias {
				if rawMessage == aliasCmd {
					h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
					if oc.Block {
						return
					}
				}
			}
			h.forMe(groupID, userID, rawMessage, atMe, oc)
		case rawMessage == "":
			h.forMe(groupID, userID, rawMessage, atMe, oc)
		}

	}
}

func (h *HandleCmd) forMe(groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand) {
	if oc.ForMe {
		h.atBotMessageHandle(groupID, userID, rawMessage, atMe, oc)
	}
}

func (h *HandleCmd) atBotMessageHandle(groupID, userID, rawMessage interface{}, atMe bool, oc global.OnCommand) {
	//fmt.Println(oc)
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

func (h *HandleCmd) runCommand(groupID, userID, rawMessage interface{}, oc global.OnCommand) {
	sendInfo := global.SendMessage{
		UserId:      userID,
		Message:     rawMessage.(string),
		AtSender:    false,
		MessageType: "private",
	}
	if groupID != nil {
		sendInfo.IsGroup = true
		sendInfo.GroupId = groupID
		sendInfo.MessageType = "group"
	}
	oc.CmdFunc(sendInfo)
}
