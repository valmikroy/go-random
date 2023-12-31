package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valmikroy/go-random/cmd/cli/command"
	"github.com/valmikroy/go-random/internal/interfaces"
)

// Reference https://github.com/PacktPublishing/Building-Modern-CLI-Applications-in-Go/tree/main/Chapter03/audiofile

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewAddCommand(client),
	}

	e := command.NewExec(cmds)
	if err := e.Execute(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
