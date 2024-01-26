package main

import (
	"github.com/NethermindEth/juno/core/felt"
	"github.com/davecgh/go-spew/spew"
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

	// declare TX
	trace, err := vm.HandleDeclareTx(declareTx, class)
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

	// deployAccount TX
	trace, err = vm.DeployAccount(declareTx.ClassHash, &felt.Zero)
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

}
