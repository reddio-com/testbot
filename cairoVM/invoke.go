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

func NewInvoke(contract string) (*core.InvokeTransaction, error) {
	InvokeTx := rpc.InvokeTxnV1{
		Version: rpc.TransactionV1,
		Type:    rpc.TransactionType_Invoke,
	}

	// Converting the contractAddress from hex to felt
	contractAddress, err := utils.HexToFelt(contract)
	if err != nil {
		panic(err.Error())
	}

	randata, err := utils.HexToFelt("0x9184e72a000")
	if err != nil {
		panic(err.Error())
	}
	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                               //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
		Calldata:           []*felt.Felt{randata},                         //this is the data that we want to pass to the function
	}

	InvokeTx.Calldata = account.FmtCalldataCairo2([]rpc.FunctionCall{FnCall})

	tx := core.InvokeTransaction{
		Nonce:              &felt.Zero,
		MaxFee:             &felt.Zero,
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		CallData:           InvokeTx.Calldata,
	}

	return &tx, nil

}
