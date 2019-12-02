package api

type LuaType = int

type ArithOp = int 		// 新增的类型别名
type CompareOp = int 	// 新增的类型别名

//定义 LuaState 接口
type LuaState interface {
	/* 基本堆栈操作 */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)

	/* 出栈方法（stack -> GO） */
	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)

	/* 入栈方法（Go -> stack） */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	// 支持运算符的4个方法
	Arith(op ArithOp)		// 执行算术和按位运算 使用运算码 区分具体执行的运算 定义在 api/consts.go中
	Compare(op CompareOp)	// 比较运算
	Len(idx int)			// 取长度运算
	Concat(n int)			// 字符串拼接
}