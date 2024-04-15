package main

import (
	"log"

	"github.com/mtaimoor1/go-dfs/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3001")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
