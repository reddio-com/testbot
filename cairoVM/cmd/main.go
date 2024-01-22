package main

import (
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

	trace, err := vm.HandleDeclareTx(declareTx, class)
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)
}
