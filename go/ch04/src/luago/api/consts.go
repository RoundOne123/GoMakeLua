package api

// 定义了一些常量
const (
	LUA_TNONE = iota - 1 // -1  对应于  Lua栈中无效索引 索引到的值
	LUA_TNIL
	LUA_TBOOLEAN
	LUA_TLIGHTUSERDATA
	LUA_TNUMBER
	LUA_TSTRING
	LUA_TTABLE
	LUA_TFUNCTION
	LUA_TUSERDATA
	LUA_TTHREAD
)