package vm

import (
	"fmt"
)

var Romsize Address

const Ramsize = 0x4000

type Handler func(s *State, i *Instruction) error

type VMErr struct {
	msg   string
	stage string
}

func (err VMErr) Error() string {
	return fmt.Sprintf("Virtual Machine %s Error: %s\n", err.stage, err.msg)
}
