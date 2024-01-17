package main

import (
	"testbot/cairoVM"
)

func main() {
	_, err := cairoVM.NewCairoVM(cairoVM.DefaultCfg())
	if err != nil {
		panic(err)
	}
}
