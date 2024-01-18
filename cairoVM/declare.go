package cairoVM

import (
	"encoding/json"
	"os"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
)

type Declare struct {
	// Fields
	Compile   string
	Sierra    string
	CasmClass *contracts.CasmClass
}

func NewDeclare(sierra_file_name string) *felt.Felt {
	// ref to https://github.com/NethermindEth/starknet.go/blob/915109ab5bc1c9c5bae7a71553a96e6665c0dcb2/account/account_test.go#L1116
	content, err := os.ReadFile(sierra_file_name)
	if err != nil {
		panic(err)
	}

	var casmClass contracts.CasmClass
	err = json.Unmarshal(content, &casmClass)
	if err != nil {
		panic(err)
	}

	compClassHash := hash.CompiledClassHash(casmClass)
	return compClassHash
}
