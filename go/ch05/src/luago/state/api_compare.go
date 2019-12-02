
package state

import . "luago/api"

// 比较运算 比较指定索引处得两个值  并返回结果  不改变栈的状态
func (self *luaState) Compare(idx1, idx2 int, op CompareOp) bool {
	a := self.stack.get(idx1)
	b := self.stack.get(idx2)
	switch op {
	case LUA_OPEQ: return _eq(a, b)
	case LUA_OPLT: return _lt(a, b)
	case LUA_OPLE: return _le(a, b)
	default: panic("invalid compare op!")
	}
}

// 比较两个值是否相等
// 只有当两个操作数再Lua语言层面具有相同类型时 才有可能返回true
// 整数、浮点数仅仅在Lua实现层面有差别 语言表层统一为 数字类型，因此需要互相转换
// 其它类型的值 按照引用进行比较
func _eq(a, b luaValue) bool {
	switch x := a.(type){
	case nil: 
		return b == nil
	case bool:
		y, ok := b.(bool)
		return ok && x == y
	case string:
		y, ok := b.(string)
		return ok, && x == y
	case int64:
		switch y := b.(type) {
		case int64: 	return x == y
		case float64: 	return float64(x) == y
		default:		return false
		}
	case float64:
		switch y := b.(type) {
		case float64: 	return x == y
		case int64:		return x == float64(y)
		default:		return false
		}
	default:
		return a == b
	}
}

// 小于操作符仅对数字和字符串类型有意义 其它情况后面再讨论 暂时调用panic终止程序
func _lt (a, b luaValue) bool {
	switch x := a.(type) {
	case string:
		if y, ok := b.(string); ok {
			return x < y
		}
	case int64:
		switch y := b.(type) {
		case int64: 	return x < y
		case float64:	return float64(x) < y
		}
	case float64:
		switch y := b.(type) {
		case float64:	return x < y
		case int64:		return x < float64(y)
		}
	}
	panic("comparison error!")
}
