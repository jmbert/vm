package vm

import "fmt"

type Ram struct {
	mem  []Byte
	size Address
}

func NewRam(size Address) *Ram {
	var r Ram
	r.size = size
	r.mem = make([]Byte, size)
	return &r
}

func (r *Ram) PeekB(addr Address) (Byte, error) {
	if addr >= r.size {
		var err VMErr

		err.stage = "Runtime"
		err.msg = fmt.Sprintf("Tried to read from out of bounds address: %4X", addr)

		return 0, err
	} else {
		return r.mem[addr], nil
	}
}

func (r *Ram) PeekW(addr Address) (Word, error) {
	var ret Word
	b1, err := r.PeekB(addr)
	if err != nil {
		return 0, err
	}
	b2, err := r.PeekB(addr + 1)
	if err != nil {
		return 0, err
	}
	ret |= Word(b1)
	ret |= Word(b2) << 8

	return ret, nil
}

func (r *Ram) PokeB(addr Address, b Byte) error {
	if addr >= r.size {
		var err VMErr

		err.stage = "Runtime"
		err.msg = fmt.Sprintf("Tried to write to out of bounds address: %4X", addr)

		return err
	} else {
		r.mem[addr] = b
		return nil
	}
}

func (r *Ram) PokeW(addr Address, w Word) error {
	b1 := Byte(w)
	b2 := Byte(w >> 8)
	err := r.PokeB(addr, b1)
	if err != nil {
		return err
	}
	err = r.PokeB(addr+1, b2)
	if err != nil {
		return err
	}
	return nil
}
