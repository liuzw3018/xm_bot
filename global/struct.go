package global

/**
 * @Author: liu zw
 * @Date: 2021/10/18 16:52
 * @File:
 * @Description: //TODO $
 * @Version:
 */

type SendMessage struct {
	UserId      interface{} // 对方 QQ 号 ( 消息类型为 private 时需要 )
	GroupId     interface{} // 群号 ( 消息类型为 group 时需要 )
	Message     interface{} // 要发送的内容
	AtSender    bool        // 是否at发送人
	IsGroup     bool        // 发送的消息是否是群消息
	MessageType string      // 消息类型, 支持 private、group , 分别对应私聊、群组, 如不传入, 则根据传入的 *_id 参数判断
}
