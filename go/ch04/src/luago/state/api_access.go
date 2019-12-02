// 和push系列方法相反 access用于从栈里获取信息
// 基本上仅使用索引访问栈里存储的信息 不会改变栈的状态

package state

import "fmt"
import . "luago/api"

// 把给定的Lua类型转换成对应的字符串表示
func (self *luaState) TypeName(tp LuaType) string {
	switch tp{
	case LUA_TNONE: 		return "no value"
	case LUA_TNIL: 			return "nil"
	case LUA_TBOOLEAN: 		return "boolean"
	case LUA_TNUMBER: 		return "number"
	case LUA_TSTRING: 		return "string"
	case LUA_TTABLE: 		return "table"
	case LUA_TFUNCTION: 	return "function"
	case LUA_TTHREAD: 		return "thread"
	default: 				return "userdata"
	}
}

// 根据索引返回值的类型 索引无效则 返回 LUA_TNONE
func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

// IsXXX 判定是否是 指定的类型
func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

// 判断给定索引处的值是否是字符串（或是 数字 涉及到Lua类型转换的讨论）
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

// 是否是（或者可以转换为）数字类型
func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

// 是否是整数类型
func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)  // 这个val.  是什么意思？ 获取类型？
	return ok
}

// 从指定索引处取出一个布尔值 如果不是布尔值 则需进行类型转换
func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

// 从指定索引取出一个数字 如果不是数字类型 则进行转换
// 无法转换成数字类型 ->  返回0
func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

// 从指定索引取出一个数字 如果不是数字类型 则进行转换
// 无法转换成数字类型 ->  返回0 + 报告是否返回成功
func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	// switch x := val.(type) {
	// case float64: 	return x, true
	// case int64: 	return float64(x), true
	// default:		return 0, false
	// }
	return convertToFloot(val)  // 使用 lua_value.go 中定义的方法
}

// 取得整数值 不是整数类型 则进行转换 
func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	//i, ok := val.(int64)
	return convertToInteger(val) // 使用 lua_value.go 中定义的方法
}

// 从指定索引处取值 如果是字符串 返回字符串
// 如果是数字 则转换为字符串（--会修改栈--） 然后返回字符串
// 否则 返回 空字符串
// 注意 C API 中只有一个返回值 因为Go语言没有nil值 故增加 ToStringX 方法特殊处理
func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type){
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s) // 这里会修改栈
		return s, true
	default:
		return "", false
	}
}