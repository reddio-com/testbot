package cairoVM

import (
	"github.com/BurntSushi/toml"
	"github.com/NethermindEth/juno/utils"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Network             utils.Network `toml:"network"`
	TxVersion           string        `toml:"tx_version"`
	AccountAddr         string        `toml:"account_addr"`
	AccountCairoVersion int           `toml:"account_cairo_version"`
}

func LoadTomlConf(fpath string) *Config {
	cfg := new(Config)
	_, err := toml.DecodeFile(fpath, cfg)
	if err != nil {
		logrus.Panicf("load config-file(%s) error: %s ", fpath, err.Error())
	}
	return cfg
}

func DefaultCfg() *Config {
	return &Config{
		Network:             1,
		TxVersion:           "0x1",
		AccountAddr:         "",
		AccountCairoVersion: 0,
	}
}
