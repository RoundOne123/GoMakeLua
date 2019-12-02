// 定义用于表示Lua值得luaValue类型
package state

import . "luago/api"

// 使用接口 表示各种不同类型得Lua值
type luaValue interface{}

// 根据变量值返回其类型
func typeOf(val luaValue) LuaType {
	switch val.(type) {
	case nil :	   return LUA_TNIL
	case bool: 	   return LUA_TBOOLEAN
	case int64:	   return LUA_TNUMBER
	case float64:  return LUA_TNUMBER
	case string:   return LUA_TSTRING
	default:	   panic("todo!")
	}
}

// 任意值转换为bool类型的值
func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:		return false
	case bool: 		return x
	default: 		return true
	}
}

// 任意值转换为浮点数
func convertToFloat(val luaValue) (float64, bool) {
	switch x := val.(type) {
	case float64 : 		return x, true
	case int64: 		return float64(x), true
	case string: 		return number.ParseFloat(x)  // parser.go  定义的方法
	default: 			return 0, false
	}
}

// 任意值转换为整数
func convertToInteger(val luaValue) (int64, bool) {
	switch x := val.(type) {
	case int64 : 		return x, true
	case float64: 		return number.FloatToInteger(x) // math.go  定义的方法
	case string: 		return _stringToInteger(x)
	default: 			return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	// 尝试直接解析为整数
	if i, ok := number.ParseInteger(s); ok {
		return i, true
	}
	// 尝试解析为浮点数
	if f, ok := number.ParseFloat(s); ok {
		// 尝试将浮点数转换为整数
		return number.FloatToInteger(f)
	}
	// 无法解析为整数
	return 0, false
}