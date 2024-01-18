package main

import (
	"testbot/cairoVM"

	"github.com/NethermindEth/juno/core"
)

func main() {
	vm, err := cairoVM.NewCairoVM(cairoVM.DefaultCfg())
	if err != nil {
		panic(err)
	}

	tx := core.DeclareTransaction{}
	_, err = vm.HandleDeclareTx(&tx)
	if err != nil {
		panic(err)
	}
}
