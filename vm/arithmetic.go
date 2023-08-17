package vm

func add_core(s *State, add Word) {
	AVal := s.GetReg(A)

	AVal += add

	if AVal == 0 {
		flags := s.GetReg(FLAGS)
		flags |= ZEROBIT
		s.SetReg(flags, FLAGS)
	}

	s.SetReg(AVal, A)
}

func AddImmediateHandler(s *State, i *Instruction) error {
	add_core(s, i.args[0])

	return nil
}

func AddRegisterHandler(s *State, i *Instruction) error {
	add_core(s, s.GetReg(Register(i.args[0])))

	return nil
}

func AddAddrHandler(s *State, i *Instruction) error {
	add, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	add_core(s, add)

	return nil
}

func sub_core(s *State, sub Word) {
	AVal := s.GetReg(A)

	AVal -= sub

	if AVal == 0 {
		flags := s.GetReg(FLAGS)
		flags |= ZEROBIT
		s.SetReg(flags, FLAGS)
	}

	s.SetReg(AVal, A)
}

func SubImmediateHandler(s *State, i *Instruction) error {
	sub_core(s, i.args[0])

	return nil
}

func SubRegisterHandler(s *State, i *Instruction) error {
	sub_core(s, s.GetReg(Register(i.args[0])))

	return nil
}

func SubAddrHandler(s *State, i *Instruction) error {
	add, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	sub_core(s, add)

	return nil
}

func mul_core(s *State, sub Word) {
	AVal := s.GetReg(A)

	AVal *= sub

	if AVal == 0 {
		flags := s.GetReg(FLAGS)
		flags |= ZEROBIT
		s.SetReg(flags, FLAGS)
	}

	s.SetReg(AVal, A)
}

func MulImmediateHandler(s *State, i *Instruction) error {
	mul_core(s, i.args[0])

	return nil
}

func MulRegisterHandler(s *State, i *Instruction) error {
	mul_core(s, s.GetReg(Register(i.args[0])))

	return nil
}

func MulAddrHandler(s *State, i *Instruction) error {
	add, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	mul_core(s, add)

	return nil
}

func div_core(s *State, sub Word) {
	AVal := s.GetReg(A)

	AVal *= sub

	if AVal == 0 {
		flags := s.GetReg(FLAGS)
		flags |= ZEROBIT
		s.SetReg(flags, FLAGS)
	}

	s.SetReg(AVal, A)
}

func DivImmediateHandler(s *State, i *Instruction) error {
	div_core(s, i.args[0])

	return nil
}

func DivRegisterHandler(s *State, i *Instruction) error {
	div_core(s, s.GetReg(Register(i.args[0])))

	return nil
}

func DivAddrHandler(s *State, i *Instruction) error {
	add, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	div_core(s, add)

	return nil
}

func IncHandler(s *State, i *Instruction) error {
	reg := s.GetReg(Register(i.args[0]))
	reg++
	s.SetReg(reg, Register(i.args[0]))
	return nil
}

func DecHandler(s *State, i *Instruction) error {
	reg := s.GetReg(Register(i.args[0]))
	reg--
	s.SetReg(reg, Register(i.args[0]))
	return nil
}
