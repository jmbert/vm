package vm

import "fmt"

func HaltHandler(s *State, i *Instruction) error {
	fmt.Printf("Halted\n")
	for {
	}
}

func NopHandler(s *State, i *Instruction) error {
	return nil
}
