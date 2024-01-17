package cairoVM

import "github.com/NethermindEth/juno/utils"

type Config struct {
	Network             utils.Network `toml:"network"`
	TxVersion           string        `toml:"tx_version"`
	AccountAddr         string        `toml:"account_addr"`
	AccountCairoVersion int           `toml:"account_cairo_version"`
}
