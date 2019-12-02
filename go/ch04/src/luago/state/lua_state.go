package state

// 定义luaState结构体 用于实现 api/lua_state.go中定义的LuaState接口
//  Go语言 不要求强制显示实现接口，只要一个结构体实现了某个接口的全部方法，就隐式实现了该接口 （相当于 只要实现了 接口中所有的方法，就相当于隐式继承了这个接口）
type luaState struct {
	stack * luaStack
}

// 创建luaState实例
func New() *luaState {
	return &luaState {
		stack: newLuaStack(20),		// 先把栈的初始容量设置为20
	}
}