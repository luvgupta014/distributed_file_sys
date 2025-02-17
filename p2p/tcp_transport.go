package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPPeer represents the remote node over TCP established connection
type TCPPeer struct {
	//conn is the underlying connection of the peer
	conn net.Conn

	//if we dial and retrieve a conn ->outbound =true
	//if we accept and retrieve a conn ->outbound =false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener //This binds the program to a specific port and starts listening for incoming client connections.

	mu   sync.RWMutex
	peer map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		//When a client connects, Accept() returns a new conn (net.Conn) object that represents the connection between the server and the client.

		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		log.Println("✅ New connection accepted from:", conn.RemoteAddr()) // Log accepted connection
		go t.handleConn(conn)

	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.Printf("New incoming connecction %+v\n", peer)
}
