// 用于定义解析方法的脚本

package number

import "strconv"

// 字符串解析为整数
func ParseInteger(str string) (int64, bool) {
	i, err := strconv.ParseInt(str, 10, 64)		// 这里直接使用了Go语言的strconv库提供的解析方法 来简化代码
	return i, err == nil
}

// 字符串解析为浮点数
func ParseFloat(str string) (float64, bool) {
	f, err := strconv.ParseFloat(str, 64)
	return f, err == nil
}