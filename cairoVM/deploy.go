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

func NewDeployCool() (*core.InvokeTransaction, error) {
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

func NewDeployInvokeTest() (*core.InvokeTransaction, error) {
	//InvokeTx := rpc.InvokeTxnV1{
	//	Version: rpc.TransactionV1,
	//	Type:    rpc.TransactionType_Invoke,
	//}

	// Converting the contractAddress from hex to felt
	// contractAddress := new(felt.Felt).SetUint64(2)
	contractAddress, _ := new(felt.Felt).SetString("0x7f2f788bcd85c25ece505a4fe359c577be77841c5afb971648af03391e5e834")

	params := new(felt.Felt).SetUint64(9099)
	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                               //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod), //this is the function that we want to call
		Calldata:           []*felt.Felt{params},                          //this is the data that we want to pass to the function
	}

	txCallData := account.FmtCallDataCairo2([]rpc.FunctionCall{FnCall})

	fmt.Println("invoke calldata = ", txCallData)

	nonce := new(felt.Felt).SetUint64(3)
	tx := core.InvokeTransaction{
		Nonce:              nonce,
		MaxFee:             &felt.Zero,
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		CallData:           txCallData,
		// CallData: []*felt.Felt{randata},
	}

	return &tx, nil

}

func NewDeployERC20() (*core.InvokeTransaction, error) {
	//InvokeTx := rpc.InvokeTxnV1{
	//	Version: rpc.TransactionV1,
	//	Type:    rpc.TransactionType_Invoke,
	//}

	// using UniversalDeploy address
	contractAddress := new(felt.Felt).SetUint64(1)

	classHash, _ := new(felt.Felt).SetString("0x781eaffc0d732cae7cf40b7488bba82d23711b7abfc06fee4c4dd2821155e18")

	salt, _ := new(felt.Felt).SetString("0x53eb1d3593b1fe9a8369a023ffa5d07d3b2050841cb75ad6ef00698d9307d10")

	uniq := new(felt.Felt).SetUint64(1)

	calldataLength := new(felt.Felt).SetUint64(4)

	name := felt.Felt{}
	name.SetString("ERC20")
	symbol := felt.Felt{}
	symbol.SetString("ERC")
	decimals := felt.Felt{}
	decimals.SetUint64(18)
	totalSupply := felt.Felt{}
	totalSupply.SetUint64(1000000000000000000)
	receipt := felt.Felt{}
	receipt.SetUint64(1)

	calldata := felt.Felt{}
	calldata.Add(&name, &symbol)

	// params := new(felt.Felt).SetUint64(8088)
	// Building the functionCall struct, where :
	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,                                                //contractAddress is the contract that we want to call
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),            //this is the function that we want to call
		Calldata:           []*felt.Felt{classHash, salt, uniq, calldataLength, &calldata}, //this is the data that we want to pass to the function
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
