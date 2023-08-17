package vm

func load_core(s *State, register Register, val Word) error {
	s.SetReg(val, register)
	return nil
}

func LoadImmediateHandler(s *State, i *Instruction) error {
	return load_core(s, Register(i.args[0]), i.args[1])
}

func LoadRegisterHandler(s *State, i *Instruction) error {
	return load_core(s, Register(i.args[0]), s.GetReg(Register(i.args[1])))
}

func LoadAddrHandler(s *State, i *Instruction) error {
	dest, err := s.r.PeekW(Address(i.args[1]))
	if err != nil {
		return err
	}
	return load_core(s, Register(i.args[0]), dest)
}

func store_core(s *State, register Register, addr Address) error {
	val := s.GetReg(register)
	return s.r.PokeW(addr, val)
}

func StoreImmediateHandler(s *State, i *Instruction) error {
	return store_core(s, Register(i.args[0]), Address(i.args[1]))
}

func StoreRegisterHandler(s *State, i *Instruction) error {
	return store_core(s, Register(i.args[0]), Address(s.GetReg(Register(i.args[1]))))
}

func StoreAddrHandler(s *State, i *Instruction) error {
	dest, err := s.r.PeekW(Address(i.args[1]))
	if err != nil {
		return err
	}
	return store_core(s, Register(i.args[0]), Address(dest))
}
