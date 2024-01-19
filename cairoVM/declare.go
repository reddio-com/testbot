package cairoVM

import (
	"encoding/json"
	"os"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
	"github.com/NethermindEth/starknet.go/rpc"
)

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
