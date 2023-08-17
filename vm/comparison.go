package vm

func compare_core(s *State, comparison Word) error {
	a := s.GetReg(A)
	if a == comparison {
		flags := s.GetReg(FLAGS)
		flags |= EQUALSBIT
		s.SetReg(flags, FLAGS)
	} else {
		flags := s.GetReg(FLAGS)
		flags &= ^(Word(EQUALSBIT))
		s.SetReg(flags, FLAGS)
	}
	return nil
}

func CompareImmediateHandler(s *State, i *Instruction) error {
	return compare_core(s, i.args[0])
}

func CompareRegisterHandler(s *State, i *Instruction) error {
	return compare_core(s, s.GetReg(Register(i.args[0])))
}

func CompareAddrHandler(s *State, i *Instruction) error {
	comparison, err := s.r.PeekW(Address(i.args[0]))
	if err != nil {
		return err
	}
	return compare_core(s, comparison)
}
