package vm

func push_core(s *State, value Word) error {
	sp := s.GetReg(SP)

	sp -= 2
	err := s.r.PokeW(Address(sp), value)
	if err != nil {
		return err
	}
	s.SetReg(sp, SP)
	return nil
}

func PushImmediateHandler(s *State, i *Instruction) error {
	return push_core(s, i.args[0])
}

func PushRegisterHandler(s *State, i *Instruction) error {
	return push_core(s, s.GetReg(Register(i.args[0])))
}

func pop_core(s *State, reg Register) error {
	sp := s.GetReg(SP)

	val, err := s.r.PeekW(Address(sp))
	sp += 2
	if err != nil {
		return err
	}
	s.SetReg(val, reg)
	s.SetReg(sp, SP)
	return nil
}

func PopHandler(s *State, i *Instruction) error {
	return pop_core(s, Register(i.args[0]))
}

func call_core(s *State, dest Word) error {
	err := push_core(s, s.GetReg(IP))
	if err != nil {
		return err
	}
	err = jump_core(s, dest)
	if err != nil {
		return err
	}
	return nil
}

func CallImmediateHandler(s *State, i *Instruction) error {
	return call_core(s, i.args[0])
}

func CallRegisterHandler(s *State, i *Instruction) error {
	return call_core(s, s.GetReg(Register(i.args[0])))
}

func RetHandler(s *State, i *Instruction) error {
	err := pop_core(s, IP)
	return err
}
