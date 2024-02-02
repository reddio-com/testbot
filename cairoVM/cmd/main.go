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
		"data/cool.sierra.json",
		"data/cool.casm.json",
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

	callClassHash, err := new(felt.Felt).SetString(cairoVM.CoolContractClassHash)
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
	deployTx, err := cairoVM.NewDeployCool()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(deployTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

	invokeTx, err = cairoVM.NewDeployInvokeTest()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(invokeTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}

	new_address, _ := new(felt.Felt).SetString("0x7f2f788bcd85c25ece505a4fe359c577be77841c5afb971648af03391e5e834")

	resp, err = vm.HandleCall(&rpc.FunctionCall{
		ContractAddress:    *new_address,
		EntryPointSelector: *utils.GetSelectorFromNameFelt("get_value"),
	},
		callClassHash,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("call response", utils.FeltToBigInt(resp[0]))

	deployTx, err = cairoVM.NewDeployCoolOld()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(deployTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}
	spew.Dump(trace)

	invokeTx, err = cairoVM.NewDeployInvokeTestCoolOld()
	if err != nil {
		panic(err)
	}

	trace, err = vm.HandleInvokeTx(invokeTx) // Assuming there is a HandleInvokeTx function
	if err != nil {
		panic(err)
	}

	new_address, _ = new(felt.Felt).SetString("0x77fcc62a59a2160f099493fcd0466c526120320c164a62a72c6ac9931db34d9")

	resp, err = vm.HandleCall(&rpc.FunctionCall{
		ContractAddress:    *new_address,
		EntryPointSelector: *utils.GetSelectorFromNameFelt("get_value"),
	},
		callClassHash,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("call response", utils.FeltToBigInt(resp[0]))

}
