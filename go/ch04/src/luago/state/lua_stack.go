package state

type luaStack struct {
	slots []luaValue	// 用来存放值
	top int 			// 记录栈顶索引
}

// 用于创建指定容量得栈
func newLuaStack(size int) *luaStack {
	return &luaStack {
		slots: make([]luaValue, size),
		top: 0,
	}
}

/*
	下面对栈的操作的函数 使用的都是绝对索引  所以对于相对索引的的栈 要进行转化
*/


// 检查栈的空闲空间是否还可以推入至少n个值 如果不能则对栈进行扩容
func (self *luaStack) check(n int) {
	free := len(self.slots) - self.top  //free 栈中空着的位置的个数
	for i := free; i < n; i++ { 		//  这里看不懂 这里像是 总共能推进 n个 ？？？？？？？
		self.slots = append(self.slots, nil)
	}
}

// 将值推入栈顶，如果溢出，则暂时调用内置的panic()函数被终止程序
func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots){
		panic("stack overflow!")
	}
	self.slots[self.top] = val 			// 直接赋值的话 栈顶 本来的值呢？ 还是说 top 表示的位置 是实际存在值得位置更上面一个索引 本来就是没值的？？？？？？
	self.top++
}

// 从栈顶弹出一个值， 如果栈是空的 则调用panic()函数终止程序
func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--						// 从这里的取值操作看 确实 top - 1 处对应的才是真正的栈顶元素  ？？？？？？
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

// 把索引转成绝对索引 未考虑索引是否有效
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	return idx + self.top + 1		// 这里是把top按照绝对索引算的 ？？？？？？ 怎么转换的？
}

// 判定索引是否有效
func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

// 根据索引从栈里取值 索引无效返回nil
func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

// 根据索引往栈里写入值 索引无效时 调用panic()函数终止程序
func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

// 将 [from, to] 区间的元素 调换位置 ？
func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}