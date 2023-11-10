package main

import (
	"flag"
	"fmt"

	greet "github.com/valmikroy/go-random/greetings"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8000, "Port number")
	flag.Parse()
	fmt.Printf("Connectin to port %d\n", port)
	fmt.Printf(greet.Hello("Blah"))
}
