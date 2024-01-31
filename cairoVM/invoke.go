package cairoVM

import (
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
)

func NewInvoke(contract string) (rpc.InvokeTxnV1, error) {
	InvokeTx := rpc.InvokeTxnV1{
		Version: rpc.TransactionV1,
		Type:    rpc.TransactionType_Invoke,
	}

	// Converting the contractAddress from hex to felt
	contractAddress, err := utils.HexToFelt(contract)
	if err != nil {
		panic(err.Error())
	}

	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress: contractAddress, //contractAddress is the contract that we want to call
		// 	EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
	}

	InvokeTx.Calldata = account.FmtCalldataCairo2([]rpc.FunctionCall{FnCall})

	return InvokeTx, FnCall, nil

}
