package main

import (
	"errors"
	goIpfs "github.com/ipfs/go-ipfs-api"

	"log"
)

var ErrCannotConnetToNode = errors.New("Check if your node(Ipfs daemon) is running. Cannot connect with node")
var ErrCidNotPinned = errors.New("The cid specified isn't pinned")

func main() {
	ipfs := goIpfs.NewLocalShell()
	if ipfs == nil {
		log.Fatal(ErrCannotConnetToNode)
	}
	path := "bafybeif2ewg3nqa33mjokpxii36jj2ywfqjpy3urdh7v6vqyfjoocvgy3a" // cid, it could be a directory or a file. Pin, Prevents the garbage collection from deleting this
	err := ipfs.Pin(path)
	if err != nil {
		log.Fatal(err)
	}

	// view pins
	pins, err := ipfs.Pins()
	if err != nil {
		log.Println(err)
	}

	if pins[path].Type == "" {
		log.Println(ErrCidNotPinned)
	} else {
		log.Printf("%s Pin-Type: %s", path, pins[path].Type)
	}

}
