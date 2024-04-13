package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listerner     net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listerner, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
}
