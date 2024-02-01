package main

import (
	"fmt"
	"testbot/cairoVM"

	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/davecgh/go-spew/spew"

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
	spew.Dump(trace)

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
	spew.Dump(trace)

	invokeTx, err := cairoVM.NewInvoke()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(invokeTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}

	callClassHash, err := new(felt.Felt).SetString("0x35eb1d3593b1fe9a8369a023ffa5d07d3b2050841cb75ad6ef00698d9307d10")
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

	resp, err := vm.HandleCall(&rpc.FunctionCall{
		ContractAddress:    *new(felt.Felt).SetUint64(2),
		EntryPointSelector: *utils.GetSelectorFromNameFelt("get_value"),
	},
		callClassHash,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("call response", utils.FeltToBigInt(resp[0]))

	// deployContract TX
	deployTx, err := cairoVM.NewDeployERC20()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(deployTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

}
