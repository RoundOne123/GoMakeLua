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

func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:		return false
	case bool: 		return x
	default: 		return true
	}
}