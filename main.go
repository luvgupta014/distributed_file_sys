package main

import (
	"fmt"
	"log"

	"github.com/anthdm/foreverstone/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
	fmt.Println("Hello, Distributed File System!")
}
