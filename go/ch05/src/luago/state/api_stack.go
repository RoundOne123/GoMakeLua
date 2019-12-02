//实现LuaState 接口中的 基础栈操纵方法

package state

// 返回栈顶索引
func (self *luaState) GetTop() int {
	return self.stack.top
}

// 把索引转换为绝对索引
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

// 检查栈的容量是否还能推入 n 个元素
// 如果不能则进行扩容 暂时忽略扩容失败的情况
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true // 忽略扩容失败的情况
}

// 从栈顶弹出n个值
func (self *luaState) Pop(n int) {
	//self.SetTop(-n-1)
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

// 把值从一个位置复制到另一个位置
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

// 把指定索引处的值推入栈顶
// 获取指定idx的值，并推入栈顶 idx的值不变
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

// 将栈顶值弹出，然后写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

// 将栈顶值弹出，然后插入指定位置
// 插入操作只是旋转操作的一种特例
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

// 删除指定索引处的值，然后将上面所有的值往下移一位
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

// 将[idx, top] 索引区间内的值 朝栈顶方向旋转n个位置
// n为负 则效果是朝栈底方向旋转
// 这里的旋转 更像是 向某个方向位移一个位置 如果唯一之后的位置不是栈内有效位置 则从反方向 再找
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	// 取得中间idx
	if n >= 0 {
		m = t - n
	}else{
		m = p - n - 1
	}
	// 调换三次区间的位置
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}

// 将栈定索引设置为指定值，
// 如果指定值小于 当前栈顶索引，则将大于指定值的元素弹出
// 如果指定值大于 当前栈顶索引，则推入多个nil值
// 指定0 相当于 清空栈
// Pop() 方法只是SetTop() 方法的特例
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx) // 获取绝对索引 作为新的top
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := self.stack.top - newTop
	// 新索引 小于 当前索引
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	}else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}