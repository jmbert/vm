package vm

type Register uint8

const (
	/* General */

	A = iota
	B
	C
	D

	X
	Y

	/* Special */

	SP
	BP

	IRT

	IP

	FLAGS

	REGNUM = iota - 1
)

type InstructionType struct {
	handler Handler
	numArgs int
}

type Opcode uint8

var opmap = map[Opcode]InstructionType{
	ADDIM:   {AddImmediateHandler, 1},
	ADDREG:  {AddRegisterHandler, 1},
	ADDADDR: {AddAddrHandler, 1},
	SUBIM:   {SubImmediateHandler, 1},
	SUBREG:  {SubRegisterHandler, 1},
	SUBADDR: {SubAddrHandler, 1},
	MULIM:   {MulImmediateHandler, 1},
	MULREG:  {MulRegisterHandler, 1},
	MULADDR: {MulAddrHandler, 1},
	DIVIM:   {DivImmediateHandler, 1},
	DIVREG:  {DivRegisterHandler, 1},
	DIVADDR: {DivAddrHandler, 1},
	INC:     {IncHandler, 1},
	DEC:     {DecHandler, 1},
	JMPIM:   {JumpImmediateHandler, 1},
	JMPREG:  {JumpRegisterHandler, 1},
	JMPADDR: {JumpAddrHandler, 1},
	JCIM:    {JumpConditionalImmediateHandler, 2},
	JCREG:   {JumpConditionalRegisterHandler, 2},
	JCADDR:  {JumpConditionalAddrHandler, 2},
	INTIM:   {InterruptImmediateHandler, 1},
	INTREG:  {InterruptRegisterHandler, 1},
	LDIM:    {LoadImmediateHandler, 2},
	LDREG:   {LoadRegisterHandler, 2},
	LDADDR:  {LoadAddrHandler, 2},
	STIM:    {StoreImmediateHandler, 2},
	STREG:   {StoreRegisterHandler, 2},
	STADDR:  {StoreAddrHandler, 2},
	HALT:    {HaltHandler, 0},
	NOOP:    {NopHandler, 0},
	CMPIM:   {CompareImmediateHandler, 1},
	CMPREG:  {CompareRegisterHandler, 1},
	CMPADDR: {CompareAddrHandler, 1},
	PUSHIM:  {PushImmediateHandler, 1},
	PUSHREG: {PushRegisterHandler, 1},
	POP:     {PopHandler, 1},
	CALLIM:  {CallImmediateHandler, 1},
	CALLREG: {CallRegisterHandler, 1},
	RET:     {RetHandler, 0},
}

/* Opcodes */

const (
	ADDIM = iota /* 0x0Y ops are Arithmetic operations : All act on A as first operand and destination */
	ADDREG
	ADDADDR

	SUBIM
	SUBREG
	SUBADDR

	MULIM
	MULREG
	MULADDR

	DIVIM
	DIVREG
	DIVADDR

	INC
	DEC
	_
	_

	JMPIM /* 0x1Y ops are Control operations */
	JMPREG
	JMPADDR

	JCIM
	JCREG
	JCADDR

	INTIM
	INTREG

	_
	_
	_
	_
	_
	_
	_
	_

	LDIM /* 0x2Y ops modify memory */
	LDREG
	LDADDR

	STIM
	STREG
	STADDR

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_

	HALT /* 0x3Y ops are miscellaneous operations */
	NOOP

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_

	CMPIM /* 0x4Y ops are comparison operations : All compare to A */
	CMPREG
	CMPADDR

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	PUSHIM /* 0x5Y ops are stack operations */
	PUSHREG

	POP

	CALLIM
	CALLREG
	RET
)

/* Conditionals */

const (
	EQUALS = iota
	NOTEQUALS
	ZERO
	NOTZERO
)

const (
	EQUALSBIT = 1
	ZEROBIT   = 2
)
