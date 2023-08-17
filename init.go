package main

import (
	"kvm/vm"

	"github.com/veandco/go-sdl2/sdl"
)

func init_machine(size vm.Address) (vm.State, *sdl.Window, *sdl.Renderer, error) {
	var state vm.State

	state.New(size)

	sdl.Init(sdl.INIT_EVERYTHING)

	win, ren, err := sdl.CreateWindowAndRenderer(vm.Width, vm.Height, 0)
	ren.SetDrawColor(0, 0, 0, 0)
	ren.Clear()
	if err != nil {
		return vm.State{}, nil, nil, err
	}

	return state, win, ren, nil
}
