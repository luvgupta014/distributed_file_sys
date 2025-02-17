package main

import (
	"distributed_file_sys/p2p"
	"log"
)

func main() {
	log.Println("Server is starting...")

	tr := p2p.NewTCPTransport(":3000") // Start server on port 4001
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	// Block main from exiting
	select {}
}
