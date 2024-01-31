package cairoVM

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
)

var (
	contractMethod string = "set_value"
)

func NewInvoke() (*core.InvokeTransaction, error) {
	InvokeTx := rpc.InvokeTxnV1{
		Version: rpc.TransactionV1,
		Type:    rpc.TransactionType_Invoke,
	}

	// Converting the contractAddress from hex to felt
	contractAddress := new(felt.Felt).SetUint64(2)

	randata := new(felt.Felt).SetUint64(2)
	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                               //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
		Calldata:           []*felt.Felt{randata},                         //this is the data that we want to pass to the function
	}

	InvokeTx.Calldata = account.FmtCalldataCairo2([]rpc.FunctionCall{FnCall})

	nonce := new(felt.Felt).SetUint64(1)
	tx := core.InvokeTransaction{
		Nonce:              nonce,
		MaxFee:             &felt.Zero,
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		CallData:           []*felt.Felt{randata},
	}

	return &tx, nil

}
