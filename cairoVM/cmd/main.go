package main

import (
	"testbot/cairoVM"

	"github.com/NethermindEth/juno/core"
)

func main() {
	vm, err := cairoVM.NewCairoVM(1)
	if err != nil {
		panic(err)
	}

	tx := core.DeclareTransaction{}
	_, err = vm.HandleDeclareTx(&tx)
	if err != nil {
		panic(err)
	}
}
