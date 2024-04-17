package p2p

type HandeshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
