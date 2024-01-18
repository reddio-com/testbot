package cairoVM

import (
	"context"
	"time"

	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/db/pebble"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
)

type Cairo struct {
	vm    vm.VM
	state core.StateReader
	acc   *Account
	cfg   *Config
}

func NewCairoVM(cfg *Config) (*Cairo, error) {
	log, err := utils.NewZapLogger(utils.ERROR, true)
	if err != nil {
		return nil, err
	}
	db, err := pebble.NewMem()
	if err != nil {
		return nil, err
	}
	txn, err := db.NewTransaction(true)
	if err != nil {
		return nil, err
	}
	return &Cairo{
		vm:    vm.New(log),
		state: core.NewState(txn),
		acc:   NewAccount(),
		cfg:   cfg,
	}, nil
}

func (c *Cairo) HandleCall(call *rpc.FunctionCall, classHash *felt.Felt) ([]*felt.Felt, error) {
	return c.vm.Call(&call.ContractAddress, classHash, &call.EntryPointSelector, call.Calldata, 0, uint64(time.Now().Unix()), c.state, c.cfg.Network)
}

func (c *Cairo) HandleDeployAccountTx(tx *core.DeployAccountTransaction) (*felt.Felt, error) {
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
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return &traces[0].ConstructorInvocation.CallerAddress, nil
}

func (c *Cairo) HandleDeclareTx(tx *core.DeclareTransaction) (*felt.Felt, error) {
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
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return traces[0].ExecuteInvocation.FunctionInvocation.ClassHash, nil
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
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.cfg.Network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
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