package models

import (
	"bytes"
	"encoding/json"
)

type Echo struct {
	Id         string `json:"Id"`
	EchoIntOne int    `json:"EchoIntOne"`
	EchoIntTwo int    `json:"EchoIntTwo"`
	EchoIntSum int    `json:"EchoIntSum"`
}

func (e *Echo) JSON() (string, error) {
	echoJSON, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(echoJSON), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
