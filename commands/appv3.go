package commands

import "fmt"

type AppCmdv3 struct {
	ID string `long:"id"`
}

func (cmd AppCmdv3) Execute(args []string) error {
	fmt.Println("CMD V3")
	return nil
}
