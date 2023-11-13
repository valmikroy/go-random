package command

import (
	"fmt"

	"github.com/valmikroy/go-random/internal/interfaces"
)

type Exec struct {
	commands []interfaces.Command
}

func NewExec(commands []interfaces.Command) *Exec {
	return &Exec{commands: commands}
}

func (e *Exec) Execute(args []string) error {

	if len(args) < 1 {
		fmt.Println("No arguments given")
		return nil
	}

	subcommand := args[0]
	for _, cmd := range e.commands {
		if cmd.Name() == subcommand {
			cmd.ParseFlags(args[1:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknow subcommand %s", subcommand)
}
