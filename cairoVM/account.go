package cairoVM

import (
	"context"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/utils"
)

type Account struct {
	ks     account.Keystore
	pubkey string
}

func NewAccount() *Account {
	ks, pub, _ := account.GetRandomKeys()
	return &Account{
		ks:     ks,
		pubkey: pub.String(),
	}
}

func (ac *Account) Sign(ctx context.Context, msg *felt.Felt) ([]*felt.Felt, error) {
	msgBig := utils.FeltToBigInt(msg)

	s1, s2, err := ac.ks.Sign(ctx, ac.pubkey, msgBig)
	if err != nil {
		return nil, err
	}
	s1Felt := utils.BigIntToFelt(s1)
	s2Felt := utils.BigIntToFelt(s2)

	return []*felt.Felt{s1Felt, s2Felt}, nil
}
