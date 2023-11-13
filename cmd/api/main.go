package main

import (
	"flag"
	"fmt"

	"github.com/valmikroy/go-random/service/echo"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 9999, "Port number")
	flag.Parse()
	fmt.Printf("Starting API at http://localhost:%d\n", port)
	echo.Run(port)
}