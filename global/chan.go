package global

/**
 * @Author: liu zw
 * @Date: 2021/10/20 11:27
 * @File:
 * @Description: 定义所需的channel
 * @Version:
 */

var (
	ErrorLogMsgChan   = make(chan string, 20) // 错误日志channel
	BotInfoLogMsgChan = make(chan string, 50) // info日志channel
)
