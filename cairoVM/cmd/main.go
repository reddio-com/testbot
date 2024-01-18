package main

import (
	"testbot/cairoVM"
	// "github.com/NethermindEth/juno/core"
)

func main() {
	_, err := cairoVM.NewCairoVM(cairoVM.DefaultCfg())
	if err != nil {
		panic(err)
	}

	declare := cairoVM.NewDeclare("data/cool_compiled_class.casm", "data/cool_sierra_contract_class.json")
	print(declare.Sierra)

	// tx := core.DeclareTransaction{}
	// _, err = vm.HandleDeclareTx(&tx)
	// if err != nil {
	// 	panic(err)
	// }
}
