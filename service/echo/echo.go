package echo

import (
	"strings"

	"github.com/valmikroy/go-random/models"

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

type EchoService struct {
	Server *http.Server
}

func CreateEchoService(port int) *EchoService {
	mux := http.NewServeMux()
	echoService := &EchoService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
	}
	mux.HandleFunc("/add", echoService.Add)
	mux.HandleFunc("/hello", echoService.Hello)
	return echoService
}

func Run(port int) *http.Server {
	service := CreateEchoService(port)
	err := service.Server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
	return service.Server
}

// Based on https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
func (e *EchoService) Add(res http.ResponseWriter, req *http.Request) {

	if req.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(req.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(res, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	req.Body = http.MaxBytesReader(res, req.Body, 1048576)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	echoObj := models.Echo{}
	err := dec.Decode(&echoObj)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {

		// Position pointer for malformed JSON
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Malformed JSON at %d", syntaxError.Offset)
			http.Error(res, msg, http.StatusBadRequest)

		// Malformed JSON
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(res, msg, http.StatusBadRequest)

		// Invalid JSON field value
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)",
				unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(res, msg, http.StatusBadRequest)

		// Unknown JSON field
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(res, msg, http.StatusBadRequest)

		// Empty request
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(res, msg, http.StatusBadRequest)

		// Large request
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(res, msg, http.StatusRequestEntityTooLarge)

		default:
			fmt.Println(err.Error())
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	echoObj.EchoIntSum = echoObj.EchoIntOne + echoObj.EchoIntTwo
	j, _ := json.Marshal(echoObj)
	io.WriteString(res, string(j))
	return
}

func (e *EchoService) Hello(res http.ResponseWriter, req *http.Request) {
	str := fmt.Sprintf("Hello World!\n")
	io.WriteString(res, str)
}
