package vm

func interrupt_call(s *State, irq Word) error {
	dest, err := s.r.PeekW(Address(irq))
	if err != nil {
		return err
	}
	err = push_core(s, s.GetReg(SP))
	if err != nil {
		return err
	}
	err = push_core(s, s.GetReg(IP))
	if err != nil {
		return err
	}
	err = jump_core(s, dest)
	if err != nil {
		return err
	}
	return nil
}

func interrupt_core(s *State, irq Word) error {
	offset := irq * 2

	irt := s.GetReg(IRT)

	ir := irt + offset

	return interrupt_call(s, ir)
}

func InterruptImmediateHandler(s *State, i *Instruction) error {
	return interrupt_core(s, i.args[0])
}

func InterruptRegisterHandler(s *State, i *Instruction) error {
	return interrupt_core(s, s.GetReg(Register(i.args[0])))
}
