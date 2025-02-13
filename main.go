package main

import (
	"distributed_file_sys/p2p"
	"fmt"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
	fmt.Println("Hello, Distributed File System!")
}
