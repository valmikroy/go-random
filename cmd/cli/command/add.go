package command

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/valmikroy/go-random/internal/interfaces"
	"github.com/valmikroy/go-random/models"
)

type AddCommand struct {
	fs         *flag.FlagSet
	client     interfaces.Client
	echoIntOne int
	echoIntTwo int
}

func NewAddCommand(client interfaces.Client) *AddCommand {
	gc := &AddCommand{
		fs:     flag.NewFlagSet("add", flag.ContinueOnError),
		client: client,
	}

	gc.fs.IntVar(&gc.echoIntOne, "one", 0, "Integer One")
	gc.fs.IntVar(&gc.echoIntTwo, "two", 0, "Integer Two")

	return gc
}

func (cmd *AddCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *AddCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: ./go-random add -one 23 -two 23")
		return fmt.Errorf("missing flags")
	}
	return cmd.fs.Parse(flags)
}

func (cmd *AddCommand) Run() error {

	echoObj := models.Echo{}
	echoObj.EchoIntOne = cmd.echoIntOne
	echoObj.EchoIntTwo = cmd.echoIntTwo

	payload, _ := json.Marshal(echoObj)
	body := bytes.NewReader(payload)

	req, err := http.NewRequest("POST", "http://localhost:9999/add", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// print the response
		fmt.Println(string(b))
	}
	return nil
}
