package binchunk

import "encoding/binary"
import "math"

type reader struct{
	data []byte 		//存放将要被解析的二进制chunk数据
}


//==========================读取基本数据类型==========================
//从字节流中读取一个字节
func (self *reader) readByte() byte{
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

//使用小端方式从字节流里读取一个cint存储类型（4个字节）
func (self *reader) readUnit32() uint32{
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

//使用小端方式从字节流里读取一个size_t存储类型（8字节）
func (self *reader) readUnit64() uint64{
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

//从字节流里读取一个Lua整数
func (self *reader) readLuaInteger() int64{
	return int64(self.readUnit64())
}

//读取一个Lua浮点数
func (self *reader) readLuaNumber() float64{
	return math.Float64frombits(self.readUnit64())
}

//读取字符串
func (self *reader) readString() string{
	size := uint(self.readByte())
	if size == 0 { 	//NULL 字符串
		return ""
	}
	if size == 0xFF { 	// 长字符串
		size = uint(self.readUnit64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes)
}

//==========================其他读取方法==========================
//读取n个字节
func (self *reader) readBytes(n uint) []byte{
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

//读取指令表
func (self *reader) readCode() []uint32 {
	code := make([]uint32, self.readUnit32())
	for i := range code {
		code[i] = self.readUnit32()
	}
	return code
}

//读取一个常量
func (self *reader) readConstant() interface{} {
	switch self.readByte() {
	case TAG_NIL:		return nil
	case TAG_BOOLEAN:	return self.readByte() != 0
	case TAG_INTEGER:	return self.readLuaInteger()
	case TAG_NUMBER:	return self.readLuaNumber()
	case TAG_SHORT_STR:	return self.readString()
	case TAG_LONG_STR:	return self.readString()
	default:			panic("corrupted!")
	}
}

//读取常量表
func (self *reader) readConstants() []interface{} {
	constants := make([]interface{}, self.readUnit32())
	for i := range constants {
		constants[i] = self.readConstant()
	}
	return constants
}

//读取UpValue表
func (self *reader) readUpvalues() []Upvalue {
	upvalues := make([]Upvalue, self.readUnit32())
	for i := range upvalues {
		upvalues[i] = Upvalue{
			Instack:	self.readByte(),
			Idx:		self.readByte(),
		}
	}
	return upvalues
}

//读取子函数原型表
func (self *reader) readProtos(parentSource string) []*Prototype {
	protos := make([]*Prototype, self.readUnit32())
	for i := range protos {
		protos[i] = self.readProto(parentSource)
	}
	return protos
}

//读取行号表
func (self *reader) readLineInfo() []uint32 {
	lineInfo := make([]uint32, self.readUnit32())
	for i := range lineInfo {
		lineInfo[i] = self.readUnit32()
	}
	return lineInfo
}

//读取局部变量表
func (self *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, self.readUnit32())
	for i := range locVars {
		locVars[i] = LocVar{
			VarName:		self.readString(),
			StartPC:		self.readUnit32(),
			EndPC:			self.readUnit32(),
		}
	}
	return locVars
}

//读取Upvalue名列表
func (self *reader) readUpvalueNames() []string {
	names := make([]string, self.readUnit32())
	for i := range names {
		names[i] = self.readString()
	}
	return names
}

//==========================检查头部==========================
func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE{
		panic("not a precompiled chunk!")
	}else if self.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	}else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	}else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	}else if self.readByte() != CINT_SIZE {
		panic("int size  mismatch!")
	}else if self.readByte() != CSZIET_SIZE {
		panic("size_t size  mismatch!")
	}else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	}else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	}else if self.readByte() != LUA_NUMBER_SIZE{
		panic("lua_Number size mismatch!")
	}else if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	}else if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}

//==========================读取函数原型==========================
func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	if source == "" {source = parentSource}
	return &Prototype{
		Source:				source,
		LineDefined:		self.readUnit32(),
		LastLineDefined:	self.readUnit32(),
		NumParams:			self.readByte(),
		IsVararg:			self.readByte(),
		MaxStackSize:		self.readByte(),
		Code:				self.readCode(),
		Constants:			self.readConstants(),
		Upvalues:			self.readUpvalues(),
		Protos:				self.readProtos(source),
		LineInfo:			self.readLineInfo(),
		LocVars:			self.readLocVars(),
		UpvalueNames:		self.readUpvalueNames(),
	}
}
