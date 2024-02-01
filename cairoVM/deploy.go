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
	deployContractMethod string = "deployContract"
)

func NewDeployERC20() (*core.InvokeTransaction, error) {
	//InvokeTx := rpc.InvokeTxnV1{
	//	Version: rpc.TransactionV1,
	//	Type:    rpc.TransactionType_Invoke,
	//}

	// using UniversalDeploy address
	contractAddress := new(felt.Felt).SetUint64(1)

	classHash, _ := new(felt.Felt).SetString("0x35eb1d3593b1fe9a8369a023ffa5d07d3b2050841cb75ad6ef00698d9307d10")

	salt, _ := new(felt.Felt).SetString("0x53eb1d3593b1fe9a8369a023ffa5d07d3b2050841cb75ad6ef00698d9307d10")

	uniq := new(felt.Felt).SetUint64(1)

	calldataLength := new(felt.Felt).SetUint64(0)

	// calldata := felt.Felt{}

	// params := new(felt.Felt).SetUint64(8088)
	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                                     //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod), //this is the function that we want to call
		Calldata:           []*felt.Felt{classHash, salt, uniq, calldataLength}, //this is the data that we want to pass to the function
	}

	txCallData := account.FmtCallDataCairo2([]rpc.FunctionCall{FnCall})

	fmt.Println("invoke calldata = ", txCallData)

	nonce := new(felt.Felt).SetUint64(2)
	tx := core.InvokeTransaction{
		Nonce:              nonce,
		MaxFee:             &felt.Zero,
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),
		CallData:           txCallData,
		// CallData: []*felt.Felt{randata},
	}

	return &tx, nil

}
