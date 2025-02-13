package main

import (
	"distributed_file_sys/p2p"
	"fmt"
	"log"
)

func main() {
	tr1 := p2p.NewTCPTransport(":4000")
	if err := tr1.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, Distributed File System!")
}
