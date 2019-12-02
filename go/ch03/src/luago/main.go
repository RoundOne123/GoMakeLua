package main

import "fmt"
import "io/ioutil"
import "os"
import "luago/binchunk"
import . "luago/vm"  //这个 . 必不可少

func main(){
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		proto := binchunk.Undump(data)
		list(proto)
	}
}

func list(f *binchunk.Prototype){
	printHeader(f) 		// 打印函数基本信息
	printCode(f)		// 打印指令表
	printDetail(f)		// 打印其他详细信息
	for _, p := range f.Protos {		// 递归调用打印子函数信息
		list(p)
	}
}

//打印函数的基本信息
func printHeader(f *binchunk.Prototype){
	funcType := "main"
	if f.LineDefined > 0 {
		funcType = "function"
	}
	varargFlag := ""
	if f.IsVararg > 0 {
		varargFlag = "+"
	}

	fmt.Printf("\n%s <%s:%d,%d> (%d instructions)\n", funcType, f.Source, f.LineDefined, f.LastLineDefined, len(f.Code))
	fmt.Printf("%d%s params, %d slots, %d upvalues, ", f.NumParams, varargFlag, f.MaxStackSize, len(f.Upvalues))
	fmt.Printf("%d locals, %d constants, %d functions\n", len(f.LocVars), len(f.Constants), len(f.Protos))
}

//打印指令表
func printCode(f *binchunk.Prototype){
	for pc, c:= range f.Code {
		line := "-"
		if len(f.LineInfo) > 0{
			line = fmt.Sprintf("%d", f.LineInfo[pc])
		}
		i := Instruction(c)
		fmt.Printf("\t%d\t[%s]\t%s \t", pc+1, line, i.OpName())
		printOperands(i)
		fmt.Printf("\n")
	}
}

//打印常量表、局部变量表和Upvalue表
func printDetail(f *binchunk.Prototype){
	fmt.Printf("constants (%d):\n", len(f.Constants))
	for i, k := range f.Constants {
		fmt.Printf("\t%d\t%s\n", i+1, constantToString(k))
	}

	fmt.Printf("locals (%d):\n", len(f.LocVars))
	for i, locVar := range f.LocVars{
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i,  locVar.VarName, locVar.StartPC+1, locVar.EndPC+1)
	}

	fmt.Printf("upvalues (%d):\n", len(f.Upvalues))
	for i, upval := range f.Upvalues {
		fmt.Printf("\t%d\t%s\t%d\t%d\n", i,  upvalName(f, i), upval.Instack, upval.Idx)
	}
}

//======== 工具方法 ========
//把常量表中的常量转字符串
func constantToString(k interface{}) string {
	switch k.(type){
	case nil:		return "nil"
	case bool:	 	return fmt.Sprintf("%t", k)
	case float64:	return fmt.Sprintf("%g", k)
	case int64:		return fmt.Sprintf("%d", k)
	case string:	return fmt.Sprintf("%q", k)
	default: return "?"
	}
}

//根据Upvalue索引从调试信息里找出Upvalue的名字
func upvalName(f *binchunk.Prototype, idx int) string {
	if len(f.UpvalueNames) > 0 {
		return f.UpvalueNames[idx]
	}
	return "-"
}

//用于打印指令的操作数
func printOperands(i Instruction) {
	switch i.OpMode() {
	case IABC:					//ABC模式
		a, b, c := i.ABC()
		fmt.Printf("%d", a)					//打印操作数A
		if i.BMode() != OpArgN {
			if b > 0xFF {
				fmt.Printf(" %d", -1-b&0xFF)
			}else{
				fmt.Printf(" %d", b)
			}
		}
		if i.CMode() != OpArgN {
			if c > 0xFF {
				fmt.Printf(" %d", -1-c&0xFF)
			}else{
				fmt.Printf(" %d", c)
			}
		}

	case IABx:					//iABx模式
		a, bx := i.ABx()
		fmt.Printf("%d", a)					//打印操作数A
		if i.BMode() == OpArgK {
			fmt.Printf(" %d", -1-bx)
		}else if i.BMode() == OpArgU {
			fmt.Printf(" %d", bx)
		}

	case IAsBx:
		a, sbx := i.AsBx()
		fmt.Printf("%d %d", a, sbx)

	case IAx:
		ax := i.Ax()
		fmt.Printf("%d", -1-ax)
	}
}