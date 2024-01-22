package cairoVM

import (
	"encoding/json"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
	"github.com/NethermindEth/starknet.go/rpc"
	"os"
)

func SetGenesis(state *core.State, sierraFileName, casmFileName string) error {

	declaredClasses := make(map[felt.Felt]core.Class)
	deployedContracts := make(map[felt.Felt]*felt.Felt)
	declaredV1Classes := make(map[felt.Felt]*felt.Felt)
	var (
		class             core.Class
		classHash         *felt.Felt
		compiledClassHash *felt.Felt
	)

	class, classHash, err := adaptClassAndHash(sierraFileName)
	if err != nil {
		return err
	}
	declaredClasses[*classHash] = class
	deployedContracts[felt.Zero] = classHash
	casmClass, err := contracts.UnmarshalCasmClass(casmFileName)
	if err != nil {
		return err
	}
	compiledClassHash = hash.CompiledClassHash(*casmClass)

	declaredV1Classes[*classHash] = compiledClassHash

	newRoot, err := new(felt.Felt).SetString("0x4427787a2736725e3f47b89f8e8a6042c9c68eda0224f0ed0b5fb414a7d79ec")
	if err != nil {
		return err
	}

	return state.Update(0, &core.StateUpdate{
		BlockHash: &felt.Zero,
		NewRoot:   newRoot,
		OldRoot:   &felt.Zero,
		StateDiff: &core.StateDiff{
			Nonces:            map[felt.Felt]*felt.Felt{felt.Zero: &felt.Zero},
			DeployedContracts: deployedContracts,
			DeclaredV1Classes: declaredV1Classes,
		},
	}, declaredClasses)
}

func NewDeclare(sierraFileName, casmFileName string) (*core.DeclareTransaction, core.Class, error) {
	// ref to https://github.com/NethermindEth/starknet.go/blob/915109ab5bc1c9c5bae7a71553a96e6665c0dcb2/account/account_test.go#L1116

	content, err := os.ReadFile(sierraFileName)
	if err != nil {
		return nil, nil, err
	}

	var class rpc.ContractClass
	err = json.Unmarshal(content, &class)
	if err != nil {
		return nil, nil, err
	}
	classHash, err := hash.ClassHash(class)

	casmClass, err := contracts.UnmarshalCasmClass(casmFileName)
	if err != nil {
		return nil, nil, err
	}

	compClassHash := hash.CompiledClassHash(*casmClass)

	var nonce felt.Felt
	nonce.SetUint64(0)

	var maxFee felt.Felt
	maxFee.SetUint64(0)

	tx := core.DeclareTransaction{
		Nonce:             &nonce,
		MaxFee:            &maxFee,
		Version:           new(core.TransactionVersion).SetUint64(2),
		CompiledClassHash: compClassHash,
		ClassHash:         classHash,
		SenderAddress:     &felt.Zero,
	}

	coreClass, err := adaptDeclaredClass(content)

	return &tx, coreClass, err
}

func adaptClassAndHash(sierraFileName string) (core.Class, *felt.Felt, error) {
	content, err := os.ReadFile(sierraFileName)
	if err != nil {
		panic(err)
	}

	var contractClass rpc.ContractClass
	err = json.Unmarshal(content, &contractClass)
	if err != nil {
		panic(err)
	}
	classHash, err := hash.ClassHash(contractClass)
	if err != nil {
		return nil, nil, err
	}
	class, err := adaptDeclaredClass(content)
	return class, classHash, err
}
