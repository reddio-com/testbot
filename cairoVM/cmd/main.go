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

	// deployAccount TX
	trace, err := vm.DeployAccount(cairoVM.AccountClassHash, &felt.Zero)
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

	// declare TX
	declareTx, class, err := cairoVM.NewDeclare(
		"data/cool_sierra_contract_class.json",
		"data/cool_compiled_class.casm",
	)
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleDeclareTx(declareTx, class)
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

}
