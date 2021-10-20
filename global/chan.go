package global

/**
 * @Author: liu zw
 * @Date: 2021/10/20 11:27
 * @File:
 * @Description: //TODO $
 * @Version:
 */

var (
	ErrorLogMsgChan   = make(chan string, 20)
	BotInfoLogMsgChan = make(chan string, 50)
)
