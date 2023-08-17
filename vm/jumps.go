package vm

import (
	"fmt"
)

func jump_core(s *State, dest Word) error {
	s.SetReg(dest, IP)
	return nil
}

func JumpImmediateHandler(s *State, i *Instruction) error {
	return jump_core(s, i.args[0])
}

func JumpRegisterHandler(s *State, i *Instruction) error {
	return jump_core(s, s.GetReg(Register(i.args[0])))
}

func JumpAddrHandler(s *State, i *Instruction) error {
	dest, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	return jump_core(s, dest)
}

func conditional_check(s *State, check Word) (bool, error) {
	var cont bool
	flags := s.GetReg(FLAGS)
	switch check {
	case EQUALS:
		if flags&EQUALSBIT != 0 {
			cont = true
		}
	case NOTEQUALS:
		if flags&EQUALSBIT == 0 {
			cont = true
		}
	case ZERO:
		if flags&ZEROBIT != 0 {
			cont = true
		}
	case NOTZERO:
		if flags&ZEROBIT == 0 {
			cont = true
		}
	default:
		var err VMErr

		err.msg = fmt.Sprintf("Unrecognised conditional %02X\n", check)
		err.stage = "Internal"

		return false, err
	}

	return cont, nil
}

func JumpConditionalImmediateHandler(s *State, i *Instruction) error {
	cont, err := conditional_check(s, i.args[0])
	if err != nil {
		return err
	} else if !cont {
		return nil
	}
	return jump_core(s, i.args[1])
}

func JumpConditionalRegisterHandler(s *State, i *Instruction) error {
	cont, err := conditional_check(s, i.args[0])
	if err != nil {
		return err
	} else if !cont {
		return nil
	}
	return jump_core(s, s.GetReg(Register(i.args[1])))
}

func JumpConditionalAddrHandler(s *State, i *Instruction) error {
	cont, err := conditional_check(s, i.args[0])
	if err != nil {
		return err
	} else if !cont {
		return nil
	}
	dest, err := s.r.PeekW(Address(i.args[1]))
	if err != nil {
		return err
	}
	return jump_core(s, dest)
}
