package global

/**
 * @Author: liu zw
 * @Date: 2021/10/19 11:26
 * @File:
 * @Description: //TODO $
 * @Version:
 */
var CmdSlice = CmdSlices{}

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
