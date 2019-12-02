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

// 用于区分Arith()方法 具体执行的运算
// 给每个算术和按位运算符都分配了一个运算码
// ？？？？？？？？？？？？？？？？疑问 如何 区分上面定义的常量 和这里的值？？？？？？？？？？？？？？？？？？
const (
	LUA_OPADD = iota 	// +
	LUA_OPSUB			// -
	LUA_OPMUL			// *
	LUA_OPMOD			// %
	LUA_OPPOW			// ^
	LUA_OPDIV 			// /
	LUA_OPIDIV 			// //
	LUA_OPBAND			// &
	LUA_OPBOR 			// |
	LUA_OPBXOR			// ~
	LUA_OPSHL 			// <<
	LUA_OPSHR 			// >>
	LUA_OPUNM 			// - (unary minus)
	LUA_OPBNOT 			// ~  不知道和上面的区别是啥？
)

// Compare() 方法对应的运算码
const (
	LUA_OPEQ = iota 	// ==
	LUA_OPLT			// <
	LUA_OPLE			// <=
)