package cairoVM

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
	"github.com/NethermindEth/starknet.go/rpc"
)

func SetGenesis(state *core.State, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	classes := make(map[felt.Felt]core.Class)
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		class, classHash, err := adaptClassAndHash(filePath)
		if err != nil {
			return err
		}
		classes[*classHash] = class
	}
	return state.Update(0, nil, classes)
}

func NewDeclare(sierraFileName string) (*core.DeclareTransaction, core.Class) {
	// ref to https://github.com/NethermindEth/starknet.go/blob/915109ab5bc1c9c5bae7a71553a96e6665c0dcb2/account/account_test.go#L1116

	content, err := os.ReadFile(sierraFileName)
	if err != nil {
		panic(err)
	}

	var class rpc.ContractClass
	err = json.Unmarshal(content, &class)
	if err != nil {
		panic(err)
	}
	classHash, err := hash.ClassHash(class)

	var casmClass contracts.CasmClass
	err = json.Unmarshal(content, &casmClass)
	if err != nil {
		panic(err)
	}

	compClassHash := hash.CompiledClassHash(casmClass)

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

	return &tx, coreClass
}

func adaptClassAndHash(fileName string) (core.Class, *felt.Felt, error) {
	content, err := os.ReadFile(fileName)
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
