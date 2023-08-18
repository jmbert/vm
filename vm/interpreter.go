package vm

import (
	"fmt"
	"os"
)

type Instruction struct {
	opcode Opcode
	args   []Word
}

func (i *Instruction) Execute(s *State) error {
	return OpMap[i.opcode].handler(s, i)
}

func DecodeFromByte(m Memory, startAddr Address) Instruction {
	var i Instruction
	opcode, err := m.PeekB(startAddr)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	var args []Word
	for index := 0; index < OpMap[Opcode(opcode)].numArgs; index++ {
		arg, err := m.PeekW(startAddr + 1 + Address(index*2))
		args = append(args, arg)
		if err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
	}

	i.opcode = Opcode(opcode)
	i.args = args

	return i
}
