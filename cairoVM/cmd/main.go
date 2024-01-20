package main

import (
	"testbot/cairoVM"
	// "github.com/NethermindEth/juno/core"
)

func main() {
	vm, err := cairoVM.NewCairoVM(cairoVM.DefaultCfg())
	if err != nil {
		panic(err)
	}

	declareTx, class, err := cairoVM.NewDeclare(
		"data/cool_sierra_contract_class.json",
		"data/cool_compiled_class.casm",
	)
	if err != nil {
		panic(err)
	}
	// fmt.Println(declare_tx)

	_, err = vm.HandleDeclareTx(declareTx, class)
	if err != nil {
		panic(err)
	}
}
