package main

import "kvm/vm"

func init_machine(size vm.Address) (vm.State, error) {
	var state vm.State

	state.New(size)

	return state, nil
}
