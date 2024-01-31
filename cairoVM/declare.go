package cairoVM

import (
	"encoding/json"
	"fmt"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
	"github.com/NethermindEth/starknet.go/rpc"
	"os"
)

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
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Declare ClassHash = ", classHash.String())

	casmClass, err := contracts.UnmarshalCasmClass(casmFileName)
	if err != nil {
		return nil, nil, err
	}

	compClassHash := hash.CompiledClassHash(*casmClass)
	fmt.Println("Declare CasmClass = ", compClassHash.String())

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
