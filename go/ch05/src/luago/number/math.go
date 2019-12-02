
package number

import "math"

// 定义整除和取模函数 为后续代码做准备
// 整除函数 （int）
func IFloorDiv(a, b int64) int64 {
	if a > 0 && b > 0 || a < 0 && b < 0 || a%b == 0 {
		return a / b
	}else {
		return a/b - 1  //向下负无穷取整
	}
}

// 整除函数 （float）
func FFloorDiv(a, b float64) float64 {
	return math.Floor(a/b)
}

// 取模函数 （int）
func IMod(a, b int64) int64 {
	return a - IFloorDiv(a, b)*b
}

// 取模函数 （int）
func FMod(a, b float64) float64 {
	return a - FFloorDiv(a, b)*b
}

// 按位 左移函数
func ShiftLeft(a, n int64) int64 {
	if n >= 0 {
		return a << uint64(n)		//Go语言里，位移运算符右边只能是无符号整数 所以要进行类型转换
	}else {
		return ShiftRight(a, -n)
	}
}

// 按位 右移函数
func ShiftRight(a, n int64) int64 {
	if n >= 0 {
		// Go语言里，右移运算符的【操作数】是有符号的，则进行的是有符号右移（空缺补1）
		// 但Lua API中是无符号右移（空缺补0） 所以要进行如下处理
		return int64(uint64(a) >> uint64(n))		
	}else {
		return ShiftLeft(a, -n)
	}
}

// 浮点数转换为整数
func FloatToInteger(f float64) (int64, bool) {
	i := int64(f)
	return i, float64(i) == f // 这样操作 效率不低吗？ 不知道
}
