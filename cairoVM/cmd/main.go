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

	declare_tx := cairoVM.NewDeclare("data/cool_sierra_contract_class.json")
	// fmt.Println(declare_tx)

	_, err = vm.HandleDeclareTx(declare_tx)
	if err != nil {
		panic(err)
	}
}
