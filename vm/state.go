package vm

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type State struct {
	r         *Ram
	registers [REGNUM + 1]Word
}

func (s *State) New(ramsize Address) {
	s.r = NewRam(ramsize)
}

func (s *State) LoadRom(r *Rom) error {
	for i := 0; i < int(r.size); i++ {
		rByte, err := r.PeekB(Address(i))
		if err != nil {
			return err
		}
		s.r.PokeB(Address(i), rByte)
	}
	return nil
}

func (s *State) SetReg(w Word, r Register) {
	s.registers[r] = w
}

func (s *State) GetReg(r Register) Word {
	return s.registers[r]
}

func (s *State) Emulate(win *sdl.Window, ren *sdl.Renderer) {
	entry, err := s.r.PeekW(Romsize - 2)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}

	s.SetReg(entry, IP)

	for {
		instr := DecodeFromByte(s.r, Address(s.GetReg(IP)))
		//fmt.Printf("Executing operation %02X at %04X\n", instr.opcode, s.GetReg(IP))
		ip := s.GetReg(IP)
		ip += 1
		ip += Word(len(instr.args) * 2)
		s.SetReg(ip, IP)
		err := instr.Execute(s)
		if err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
		err = Display(s, win, ren)
		if err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
	}
}
