package cairoVM

import (
	"fmt"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
)

var (
	contractMethod string = "set_value"
)

func SetCoolValue(address string, value uint64) (*core.InvokeTransaction, error) {
	contractAddress, err := new(felt.Felt).SetString(address)
	if err != nil {
		return nil, err
	}

	params := new(felt.Felt).SetUint64(value)

	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                               //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
		Calldata:           []*felt.Felt{params},                          //this is the data that we want to pass to the function
	}

	txCallData := account.FmtCallDataCairo2([]rpc.FunctionCall{FnCall})

	fmt.Println("invoke calldata = ", txCallData)

	tx := core.InvokeTransaction{
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		CallData:           txCallData,
	}

	return &tx, nil
}
