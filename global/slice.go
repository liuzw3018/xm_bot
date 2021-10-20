package global

/**
 * @Author: liu zw
 * @Date: 2021/10/19 11:26
 * @File:
 * @Description: 全局所需的切片定义
 * @Version:
 */
// CmdSlice 插件列表
var CmdSlice = CmdSlices{}

// CmdSlices 插件列表，可根据优先级排序
type CmdSlices []OnCommand

func (cs CmdSlices) Len() int {
	return len(cs)
}

func (cs CmdSlices) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func (cs CmdSlices) Less(i, j int) bool {
	return cs[i].Priority < cs[j].Priority
}
