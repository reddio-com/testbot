package cairoVM

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NethermindEth/juno/adapters/sn2core"
	"github.com/NethermindEth/juno/blockchain"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/db"
	"github.com/NethermindEth/juno/starknet"
)

func init() {
	blockchain.RegisterCoreTypesToEncoder()
}

var AccountClassHash *felt.Felt

func BuildGenesis(classesPaths []string) (*blockchain.PendingStateWriter, error) {
	classes, err := loadClasses(classesPaths)
	if err != nil {
		return nil, err
	}
	genesisState := blockchain.NewPendingStateWriter(core.EmptyStateDiff(), make(map[felt.Felt]core.Class), core.NewState(db.NewMemTransaction()))

	var addr uint64
	for classHash, class := range classes {
		// Sets pending.newClasses, DeclaredV0Classes, (not DeclaredV1Classes)
		if err = genesisState.SetContractClass(&classHash, class); err != nil {
			return nil, fmt.Errorf("declare class: %v", err)
		}

		if cairo1Class, isCairo1 := class.(*core.Cairo1Class); isCairo1 {
			if err = genesisState.SetCompiledClassHash(&classHash, cairo1Class.Compiled.Hash()); err != nil {
				return nil, fmt.Errorf("set compiled class hash: %v", err)
			}
		}

		addrFelt := new(felt.Felt).SetUint64(addr)
		fmt.Println("Genesis.SetClassHash = ", classHash.String())
		err = genesisState.SetClassHash(addrFelt, &classHash)
		if err != nil {
			return nil, err
		}
		addr++
	}

	return genesisState, nil
}

// return map[classHash]Class
func loadClasses(classes []string) (map[felt.Felt]core.Class, error) {
	classMap := make(map[felt.Felt]core.Class)
	for i, classPath := range classes {
		bytes, err := os.ReadFile(classPath)
		if err != nil {
			return nil, fmt.Errorf("read class file: %v", err)
		}

		var response *starknet.ClassDefinition
		if err = json.Unmarshal(bytes, &response); err != nil {
			return nil, fmt.Errorf("unmarshal class: %v", err)
		}

		var coreClass core.Class
		if response.V0 != nil {
			if coreClass, err = sn2core.AdaptCairo0Class(response.V0); err != nil {
				return nil, err
			}
		} else if compiledClass, cErr := starknet.Compile(response.V1); cErr != nil {
			return nil, cErr
		} else if coreClass, err = sn2core.AdaptCairo1Class(response.V1, compiledClass); err != nil {
			return nil, err
		}

		classhash, err := coreClass.Hash()
		if err != nil {
			return nil, fmt.Errorf("calculate class hash: %v", err)
		}
		if i == 0 {
			AccountClassHash = classhash
		}
		classMap[*classhash] = coreClass
	}
	return classMap, nil
}
