package cairoVM

import (
	"fmt"
	"os"
)

type Declare struct {
	// Fields
	Compile string
	Sierra  string
}

func NewDeclare(compile_casm_file_name string, sierra_file_name string) *Declare {
	casm_data, err := os.ReadFile(compile_casm_file_name) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	sierra_data, err := os.ReadFile(sierra_file_name) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	return &Declare{
		Compile: string(casm_data),
		Sierra:  string(sierra_data),
	}
}
