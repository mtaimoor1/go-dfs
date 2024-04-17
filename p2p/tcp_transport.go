package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represent the remote node over TCP connection
type TCPPeer struct {
	conn     net.Conn
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
	listerner     net.Listener
	handshakeFunc HandeshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
		handshakeFunc: NOPHandshakeFunc,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listerner, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listerner.Accept()
		if err != nil {
			fmt.Printf("Error while Accepting Connection: %s \n", err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming connection: %v", peer)

	// Handshake
	if err := t.handshakeFunc(peer); err != nil {
		fmt.Printf("Error while handshake: %s\n", err)
		fmt.Print("Closing TCP connection!!!")
		conn.Close()
		return
	}

	msg := &Temp{}

	// Read Loop
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Decoding Error, %s\n", err)
			continue
		}
	}
}

type Temp struct{}

// 51:05
