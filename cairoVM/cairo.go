package cairoVM

import (
	"context"
	"fmt"
	"time"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
)

type Cairo struct {
	vm    vm.VM
	state core.StateReader
	acc   *Account
	cfg   *Config

	TxVersion *core.TransactionVersion
	MaxFee    *felt.Felt
}

func NewCairoVM(cfg *Config) (*Cairo, error) {
	log, err := utils.NewZapLogger(utils.ERROR, true)
	if err != nil {
		return nil, err
	}
	//db, err := pebble.NewMem()
	//if err != nil {
	//	return nil, err
	//}
	//txn, err := db.NewTransaction(true)
	//if err != nil {
	//	return nil, err
	//}
	//state := core.NewState(txn)
	//cairoFiles := make(map[string]string)
	//cairoFiles["data/genesis/NoValidateAccount.sierra.json"] = "data/genesis/NoValidateAccount.casm.json"
	//cairoFiles["data/genesis/erc20.sierra.json"] = "data/genesis/erc20.casm.json"

	state, err := BuildGenesis(
		[]string{
			"data/genesis/NoValidateAccount.sierra.json",
			"data/genesis/UniversalDeployer.json",
			"data/genesis/cool_sierra_contract_class.json",
		},
	)
	if err != nil {
		return nil, err
	}
	return &Cairo{
		vm:        vm.New(log),
		state:     state,
		acc:       NewAccount(),
		cfg:       cfg,
		TxVersion: new(core.TransactionVersion).SetUint64(cfg.TxVersion),
		MaxFee:    new(felt.Felt).SetUint64(cfg.MaxFee),
	}, nil
}

func (c *Cairo) HandleCall(call *rpc.FunctionCall, classHash *felt.Felt) ([]*felt.Felt, error) {
	return c.vm.Call(&call.ContractAddress, classHash, &call.EntryPointSelector, call.Calldata, 0, uint64(time.Now().Unix()), c.state, c.cfg.Network)
}

func (c *Cairo) DeployAccount(classHash, contractAddr *felt.Felt) (*vm.TransactionTrace, error) {
	tx := &core.DeployAccountTransaction{
		DeployTransaction: core.DeployTransaction{
			ContractAddressSalt: c.acc.pubkey,
			ContractAddress:     contractAddr,
			ClassHash:           classHash,
			ConstructorCallData: []*felt.Felt{c.acc.pubkey},
			Version:             c.TxVersion,
		},
		MaxFee: c.MaxFee,
		Nonce:  &felt.Zero,
	}
	if tx.ContractAddress == nil {
		tx.ContractAddress = core.ContractAddress(&felt.Zero, tx.ClassHash, tx.ContractAddressSalt, tx.ConstructorCallData)
	}
	return c.HandleDeployAccountTx(tx)
}

func (c *Cairo) HandleDeployAccountTx(tx *core.DeployAccountTransaction) (*vm.TransactionTrace, error) {
	fmt.Println("------------- DeployAccount TX -------------")
	txnHash, err := core.TransactionHash(tx, c.cfg.Network)
	if err != nil {
		return nil, err
	}
	tx.TransactionHash = txnHash
	sig, err := c.acc.Sign(context.Background(), txnHash)
	if err != nil {
		return nil, err
	}
	tx.TransactionSignature = sig

	txs := []core.Transaction{tx}
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, make([]*felt.Felt, 0), false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return &traces[0], nil
}

func (c *Cairo) HandleDeclareTx(tx *core.DeclareTransaction, class core.Class) (*vm.TransactionTrace, error) {
	fmt.Println("------------- Declare TX -------------")
	txnHash, err := core.TransactionHash(tx, c.cfg.Network)
	if err != nil {
		return nil, err
	}
	tx.TransactionHash = txnHash
	sig, err := c.acc.Sign(context.Background(), txnHash)
	if err != nil {
		return nil, err
	}
	tx.TransactionSignature = sig

	txs := []core.Transaction{tx}
	classes := []core.Class{class}
	_, traces, err := c.vm.Execute(txs, classes, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, make([]*felt.Felt, 0), false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}

	return &traces[0], nil
}

func (c *Cairo) HandleInvokeTx(tx *core.InvokeTransaction) (*vm.TransactionTrace, error) {
	txnHash, err := core.TransactionHash(tx, c.cfg.Network)
	if err != nil {
		return nil, err
	}
	tx.TransactionHash = txnHash
	sig, err := c.acc.Sign(context.Background(), txnHash)
	if err != nil {
		return nil, err
	}
	tx.TransactionSignature = sig

	txs := []core.Transaction{tx}
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, make([]*felt.Felt, 0), false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return &traces[0], nil
}

func (c *Cairo) HandleL1HandlerTx(tx *core.L1HandlerTransaction) error {
	// L1 handle tx has no signature
	txnHash, err := core.TransactionHash(tx, c.cfg.Network)
	if err != nil {
		return err
	}
	tx.TransactionHash = txnHash

	txs := []core.Transaction{tx}
	_, _, err = c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, []*felt.Felt{&felt.Zero}, false, false, true, &felt.Zero, &felt.Zero, false)
	return err
}
