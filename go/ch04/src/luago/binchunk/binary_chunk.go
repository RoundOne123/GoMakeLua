package binchunk

type binChunk struct{
	header					// 头部
	sizeUpvalues byte 		// 记录主函数upvalue数量的字段
	mainFunc *Prototype		// 主函数原型
}

// 头部
type header struct{
	signature 		[4]byte  //魔数/签名
	version			byte
	format			byte
	luacData		[6]byte
	cintSize		byte
	sizetSize		byte
	instructionSize	byte
	luaIntegerSize	byte
	luaNumberSize	byte
	luacInt			int64
	luacNum			float64
}

//头部各部分 对应的值
const (
	LUA_SIGNATURE = "\x1bLua"
	LUAC_VERSION  = 0x53
	LUAC_FORMAT = 0
	LUAC_DATA = "\x19\x93\r\n\x1a\n"
	CINT_SIZE = 4
	CSZIET_SIZE = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE = 8
	LUAC_INT = 0x5678
	LUAC_NUM = 370.5
)

//函数原型
type Prototype struct{
	Source			string  			//源文件名
	LineDefined		uint32 				
	LastLineDefined	uint32
	NumParams		byte
	IsVararg		byte
	MaxStackSize	byte
	Code			[]uint32
	Constants		[]interface{}
	Upvalues  		[]Upvalue
	Protos			[]*Prototype
	LineInfo		[]uint32
	LocVars			[]LocVar
	UpvalueNames	[]string
}

//tag值常量
const (
	TAG_NIL		= 0x00
	TAG_BOOLEAN = 0x01
	TAG_NUMBER	= 0x03
	TAG_INTEGER = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR = 0x14
)

//Upvalue
type Upvalue struct{
	Instack byte
	Idx 	byte
}

//局部变量表
type LocVar struct{
	VarName string
	StartPC uint32
	EndPC	uint32
}

// 解析二进制chunk
func Undump(data []byte) *Prototype{
	reader := &reader{data}
	reader.checkHeader()			//校验头部
	reader.readByte()				//跳过UpValue数量
	return reader.readProto("")		//读取函数原型
}