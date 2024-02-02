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
	contractAddress, err := new(felt.Felt).SetString(UniversalDeployerContractAddress)

	if err != nil {
		return nil, err
	}

	classHash, err := new(felt.Felt).SetString(CoolContractClassHash)

	if err != nil {
		return nil, err
	}

	salt, err := new(felt.Felt).SetString(RandomSalt)
	if err != nil {
		return nil, err
	}

	uniq := new(felt.Felt).SetUint64(1)

	calldataLength := new(felt.Felt).SetUint64(0)

	FnCall := rpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),
		Calldata:           []*felt.Felt{classHash, salt, uniq, calldataLength},
	}

	txCallData := account.FmtCallDataCairo2([]rpc.FunctionCall{FnCall})

	nonce := new(felt.Felt).SetUint64(2)
	tx := core.InvokeTransaction{
		Nonce:              nonce,
		MaxFee:             &felt.Zero,
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),
		CallData:           txCallData,
	}
	return &tx, nil

}

func NewDeployInvokeTest() (*core.InvokeTransaction, error) {
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

	tx := core.InvokeTransaction{
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		CallData:           txCallData,
		// CallData: []*felt.Felt{randata},
	}

	return &tx, nil

}

func NewDeployCoolOld() (*core.InvokeTransaction, error) {
	contractAddress := new(felt.Felt).SetUint64(1)

	classHash, _ := new(felt.Felt).SetString("0x47f93257c3a6e42fc71162a646b3223dfad27c2d994f97f333492c66e31b8c8")

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

	tx := core.InvokeTransaction{
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),
		CallData:           txCallData,
		// CallData: []*felt.Felt{randata},
	}

	return &tx, nil

}

func NewDeployInvokeTestCoolOld() (*core.InvokeTransaction, error) {
	contractAddress, _ := new(felt.Felt).SetString("0x77fcc62a59a2160f099493fcd0466c526120320c164a62a72c6ac9931db34d9")

	params := new(felt.Felt).SetUint64(1111)
	// Building the functionCall struct, where :
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
