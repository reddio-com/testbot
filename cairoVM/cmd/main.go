package main

import (
	"testbot/cairoVM"
)

func main() {
	_, err := cairoVM.NewCairoVM(1)
	if err != nil {
		panic(err)
	}
}
