package main

import (
	"fmt"
	"testbot/cairoVM"
	// "github.com/NethermindEth/juno/core"
)

func main() {
	_, err := cairoVM.NewCairoVM(cairoVM.DefaultCfg())
	if err != nil {
		panic(err)
	}

	declare_hash := cairoVM.NewDeclare("data/cool_sierra_contract_class.json")
	fmt.Println(declare_hash)

	// tx := core.DeclareTransaction{}
	// _, err = vm.HandleDeclareTx(&tx)
	// if err != nil {
	// 	panic(err)
	// }
}
