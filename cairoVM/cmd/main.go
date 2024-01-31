package main

import (
	"fmt"
	"testbot/cairoVM"

	"github.com/NethermindEth/juno/core/felt"
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
	// spew.Dump(trace)
	fmt.Println("trace type ", trace.Type)

	// declare TX
	declareTx, class, err := cairoVM.NewDeclare(
		"data/erc20.sierra.json",
		"data/erc20.casm.json",
	)
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleDeclareTx(declareTx, class)
	if err != nil {
		panic(err)
	}
	// spew.Dump(trace)

	invokeTx, err := cairoVM.NewInvoke()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(invokeTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}
	// spew.Dump(trace)

}
