package main

import (
	"github.com/reddio-com/testbot/cairoVM"
)

func main() {
	_, err := cairoVM.NewCairoVM(1)
	if err != nil {
		panic(err)
	}
}
