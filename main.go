package main

import (
	"fmt"
	"kvm/vm"
	"os"
)

func main() {
	state, win, ren, err := init_machine(vm.Ramsize)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	romFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	rInfo, err := romFile.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	vm.Romsize = vm.Address(rInfo.Size())

	rom := vm.NewRom(vm.Romsize)

	rom.FromFile(romFile)

	state.LoadRom(rom)

	state.Emulate(win, ren)
}
