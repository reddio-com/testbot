package cairoVM

import (
	"fmt"
	"github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/contracts"
	"github.com/NethermindEth/starknet.go/hash"
)

func init() {
	blockchain.RegisterCoreTypesToEncoder()
}

func SetGenesis(state *core.State, cairoFiles map[string]string) error {

	declaredClasses := make(map[felt.Felt]core.Class)
	deployedContracts := make(map[felt.Felt]*felt.Felt)
	declaredV1Classes := make(map[felt.Felt]*felt.Felt)
	nonces := make(map[felt.Felt]*felt.Felt)
	var (
		class             core.Class
		classHash         *felt.Felt
		compiledClassHash *felt.Felt
		err               error
	)

	var addr uint64 = 0

	for sierraFileName, casmFileName := range cairoFiles {
		class, classHash, err = adaptClassAndHash(sierraFileName)
		if err != nil {
			return err
		}
		fmt.Println("genesis classHash = ", classHash.String())
		declaredClasses[*classHash] = class
		addrFelt := new(felt.Felt).SetUint64(addr)
		deployedContracts[*addrFelt] = classHash
		casmClass, err := contracts.UnmarshalCasmClass(casmFileName)
		if err != nil {
			return err
		}
		compiledClassHash = hash.CompiledClassHash(*casmClass)

		declaredV1Classes[*classHash] = compiledClassHash

		nonces[*addrFelt] = &felt.Zero
		addr++
	}

	newRoot, err := new(felt.Felt).SetString("0x56f007b0f69daa75af325ecfa0d717bfd4d72bfa102151912fe4a15b9dfd30f")
	if err != nil {
		return err
	}

	return state.Update(0, &core.StateUpdate{
		BlockHash: &felt.Zero,
		NewRoot:   newRoot,
		OldRoot:   &felt.Zero,
		StateDiff: &core.StateDiff{
			Nonces:            nonces,
			DeployedContracts: deployedContracts,
			DeclaredV1Classes: declaredV1Classes,
		},
	}, declaredClasses)
}
