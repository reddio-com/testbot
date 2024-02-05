package cairoVM

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
)

var (
	deployContractMethod string = "deployContract"
)

func NewDeployCool(contractClassHash string) (*core.InvokeTransaction, error) {
	contractAddress, err := new(felt.Felt).SetString(UniversalDeployerContractAddress)

	if err != nil {
		return nil, err
	}

	classHash, err := new(felt.Felt).SetString(contractClassHash)

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

	tx := core.InvokeTransaction{
		Version:            new(core.TransactionVersion).SetUint64(1),
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(deployContractMethod),
		CallData:           txCallData,
	}
	return &tx, nil

}
