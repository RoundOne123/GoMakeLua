package vm

//特殊常量，可以认为是一个可以被编译器修改的常量。
//在每一个const关键字出现时，被重置为0，'
//然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

const (
	OP_MOVE = iota
	OP_LOADK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota
	OpArgU
	OpArgR
	OpArgK
)

//定义opcode结构体
type opcode struct{
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode   byte
	name	 string
}

//定义完整的指令表
var opcodes = []opcode{
	/*	   T  A     B      C     mode   name     */
	opcode{0, 1, OpArgR, OpArgN, IABC, "MOVE     "},
	opcode{0, 1, OpArgK, OpArgN, IABx, "LOADK    "},
	opcode{0, 1, OpArgN, OpArgN, IABx, "LOADKX   "},
	opcode{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL "},
	opcode{0, 1, OpArgU, OpArgN, IABC, "LOADNIL  "},
	opcode{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL "},
	opcode{0, 1, OpArgU, OpArgK, IABC, "GETTABUP "},
	opcode{0, 1, OpArgR, OpArgK, IABC, "GETTABLE "},
	opcode{0, 0, OpArgK, OpArgK, IABC, "GETTABUP "},
	opcode{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL "},
	opcode{0, 0, OpArgK, OpArgK, IABx, "SETTABLE "},
	opcode{0, 1, OpArgU, OpArgU, IABx, "NEWTABLE "},
	opcode{0, 1, OpArgR, OpArgK, IABC, "SELF     "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "ADD      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SUB      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MUL      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "MOD      "},
	opcode{0, 0, OpArgK, OpArgK, IABC, "POW      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "DIV      "},
	opcode{0, 1, OpArgK, OpArgK, IABx, "IDIV     "},
	opcode{0, 1, OpArgK, OpArgK, IABx, "BAND     "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BOR      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "BXOR     "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHL      "},
	opcode{0, 1, OpArgK, OpArgK, IABC, "SHR      "},
	opcode{0, 1, OpArgR, OpArgN, IABC, "UNM      "},
	opcode{0, 1, OpArgR, OpArgN, IABC, "BNOT     "},
	opcode{0, 1, OpArgR, OpArgN, IABC, "NOT      "}, 
	opcode{0, 1, OpArgR, OpArgN, IABC, "LEN      "},
	opcode{0, 1, OpArgR, OpArgR, IABC, "CONCAT   "},
	opcode{0, 0, OpArgR, OpArgN, IAsBx, "JMP      "},
	opcode{1, 0, OpArgK, OpArgK, IABC, "EQ       "},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LT       "},
	opcode{1, 0, OpArgK, OpArgK, IABC, "LE       "},
	opcode{1, 0, OpArgN, OpArgU, IABC, "TEST     "},
	opcode{1, 1, OpArgR, OpArgU, IABx, "TESTSET  "},
	opcode{0, 1, OpArgU, OpArgU, IABx, "CALL     "},
	opcode{0, 1, OpArgU, OpArgU, IABC, "TAILCALL "},
	opcode{0, 0, OpArgU, OpArgN, IABC, "RETURN   "},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP  "},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP  "},
	opcode{0, 0, OpArgN, OpArgU, IABC, "TFORCALL "},
	opcode{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP "},
	opcode{0, 0, OpArgU, OpArgU, IABC, "SETLIST  "},
	opcode{0, 1, OpArgU, OpArgN, IABx, "CLOSURE  "},
	opcode{0, 1, OpArgU, OpArgN, IABC, "VARARG   "},
	opcode{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG  "},
}