// 把算术和位移运算符统一 映射为 Go语言运算符 或者 前面预先定义好的函数
// 从栈的角度考虑 
// 一元运算 从栈顶去除值 计算  然后把 结果存到栈顶
// 二元运算 依次从栈顶取出两个值 然后把 结果存到 栈顶

package state

import "math"
import . "luago/api"
import "luago/number"

var (
	// 这些全部 接收两个参数 返回一个参数
	iadd = func(a, b int64) int64 { return a + b }
	fadd = func(a, b float64) float64 { return a + b }
	isub = func(a, b int64) int64 { return a - b }
	fsub = func(a, b float64) float64 { return a - b }
	imul = func(a, b int64) int64 { return a * b }
	fmul = func(a, b float64) float64 { return a * b }
	imod = number.IMod
	fmod = number.FMod 
	pow  = math.Pow 
	div  = func(a, b float64) float64 { return a / b }
	iidiv = number.IFloorDiv
	didiv = number.FFloorDiv
	band = func(a, b int64) int64 { return a & b }
	bor  = func(a, b int64) int64 { return a | b }
	bxor = func(a, b int64) int64 { return a ^ b }
	shl  = number.ShiftLeft
	shr  = number.ShiftRight
	iunm = func(a, _ int64) int64 { return  - a }
	funm = func(a, _ float64) float64 { return - a }
	bnot = func(a, _ int64) int64 { return ^a }
)

// 定义一个结构体 容纳整数和浮点数类型的运算
type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc func(float64, float64) float64
}

// 各种运算 顺序要和Lua运算码常量顺序一致
var operators = []operator{
	operator{iadd, fadd},
	operator{isub, fsub},
	operator{imul, fmul},
	operator{imod, fmod},
	operator{nil, pow},
	operator{nil, div},
	operator{iidiv, fidiv},
	operator{band, nil},
	operator{bor, nil},
	operator{bxor, nil},
	operator{shl, nil},
	operator{shr, nil},
	operator{iunm, funm},
	operator{bnot, nil},
}



// 算术和按位运算
func (self *luaState) Arith(op ArithOp) {
	// 弹出一个或两个操作数
	var a, b luaValue	// operands
	b = self.stack.pop()
	if op != LUA_OPUNM && op != LUA_OPBNOT {
		a = self.stack.pop()
	}else {
		a = b
	}

	// 按照索引 取出opertor实例
	operator := operators[op]
	// 调用_arith 执行计算
	result := _arith(a, b, operator); 
	// 将结果推入栈
	if result != nil {
		self.stack.push(result)
	}else {
		panic("arithmetic error!")
	}
}

// 执行计算
func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil {	//bitwise 按位运算 操作数都是（或可转换为）整数
		if x, ok := convertToInteger(a); ok {
			if y, ok : = convertToInteger(b); ok {
				return op.integerFunc(x, y) // 这里得 integerFunc  floatFunc 定义在哪里？
			}
		}
	} else {	//arith 
		if op.integerFunc != nil { 	// add, sub, mul, mod, idiv, unm 操作数是整数时进行运算
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		// 其它情况 尝试将操作数转换为 浮点数再执行运算
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}