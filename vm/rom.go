package vm

import (
	"fmt"
	"os"
)

type Rom struct {
	mem  []Byte
	size Address
}

func NewRom(size Address) *Rom {
	var r Rom
	r.size = size
	r.mem = make([]Byte, size)
	return &r
}

func (r *Rom) FromFile(f *os.File) {
	for i := 0; ; i++ {
		var buf []byte
		buf = append(buf, 0)
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		r.mem[Address(i)] = Byte(buf[0])
	}
}

func (r *Rom) PeekB(addr Address) (Byte, error) {
	if addr >= r.size {
		var err VMErr

		err.stage = "Runtime"
		err.msg = fmt.Sprintf("Tried to read from out of bounds address: %4X", addr)

		return 0, err
	} else {
		return r.mem[addr], nil
	}
}

func (r *Rom) PeekW(addr Address) (Word, error) {
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

func (r *Rom) PokeB(addr Address, b Byte) error {
	var err VMErr

	err.stage = "Runtime"
	err.msg = fmt.Sprintf("Tried to write to ROM address: %4X", addr)

	return err
}

func (r *Rom) PokeW(addr Address, w Word) error {
	var err VMErr
	err.stage = "Runtime"
	err.msg = fmt.Sprintf("Tried to write to ROM address: %4X", addr)
	return nil
}
