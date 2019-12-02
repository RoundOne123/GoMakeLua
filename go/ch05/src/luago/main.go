package main

import "fmt"
import . "luago/api"
import "luago/state"

func main() {
	ls := state.New()
	ls.PushBoolean(true); 	printStack(ls)
	ls.PushInteger(10);		printStack(ls)
	ls.PushNil();			printStack(ls)
	ls.PushString("hello"); printStack(ls)
	ls.PushValue(-4);		printStack(ls)      // 获取指定idx的值，并推入栈顶 idx的值不变
	ls.Replace(3);			printStack(ls)		// 将栈顶值弹出，然后写入指定位置
	ls.SetTop(6);			printStack(ls)
	ls.Remove(-3);			printStack(ls)
	ls.SetTop(-5);			printStack(ls)
}

// 打印栈的内容
func printStack(ls LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case LUA_TBOOLEAN:  fmt.Printf("[%t]", ls.ToBoolean(i))
		case LUA_TNUMBER: 	fmt.Printf("[%g]", ls.ToNumber(i))
		case LUA_TSTRING:   fmt.Printf("[%q]", ls.ToString(i))
		default:			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}