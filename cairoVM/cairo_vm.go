package cairoVM

import (
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/juno/db/pebble"
	"github.com/NethermindEth/juno/rpc"
	"github.com/NethermindEth/juno/utils"
	"github.com/NethermindEth/juno/vm"
	"time"
)

type CairoVM struct {
	vm      vm.VM
	state   core.StateReader
	network utils.Network
}

func NewCairoVM(network utils.Network) (*CairoVM, error) {
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
	return &CairoVM{
		vm:      vm.New(log),
		state:   core.NewState(txn),
		network: network,
	}, nil
}

func (c *CairoVM) HandleCall(call *rpc.FunctionCall, classHash *felt.Felt) ([]*felt.Felt, error) {
	return c.vm.Call(&call.ContractAddress, classHash, &call.EntryPointSelector, call.Calldata, 0, uint64(time.Now().Unix()), c.state, c.network)
}

func (c *CairoVM) HandleDeployAccountTx(tx *core.DeployAccountTransaction) (*felt.Felt, error) {
	txs := []core.Transaction{tx}
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return &traces[0].ConstructorInvocation.CallerAddress, nil
}

func (c *CairoVM) HandleDeclareTx(tx *core.DeclareTransaction) (*felt.Felt, error) {
	txs := []core.Transaction{tx}
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return traces[0].ExecuteInvocation.FunctionInvocation.ClassHash, nil
}

func (c *CairoVM) HandleInvokeTx(tx *core.InvokeTransaction) (*vm.TransactionTrace, error) {
	txs := []core.Transaction{tx}
	_, traces, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.network, nil, false, false, true, &felt.Zero, &felt.Zero, false)
	if err != nil {
		return nil, err
	}
	return &traces[0], nil
}

func (c *CairoVM) HandleL1HandlerTx(tx *core.L1HandlerTransaction) error {
	txs := []core.Transaction{tx}
	_, _, err := c.vm.Execute(txs, nil, 0, uint64(time.Now().Unix()), &felt.Zero, c.state, c.network, []*felt.Felt{&felt.Zero}, false, false, true, &felt.Zero, &felt.Zero, false)
	return err
}
