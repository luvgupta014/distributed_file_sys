package p2p

//peer is an interface that represents the remote node
type Peer interface{}

//transport is anything that handles the communication
//between the nodes in the network. This can be tcp, udp, websockets
type Transport interface {
	ListenAndAccept() error
}
