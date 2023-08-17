package vm

type Word uint16
type Byte uint8
type Address uint16

type Memory interface {
	PeekB(addr Address) (Byte, error)
	PeekW(addr Address) (Word, error)
	PokeB(addr Address, b Byte) error
	PokeW(addr Address, b Word) error
}
