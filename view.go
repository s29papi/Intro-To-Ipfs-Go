package main

import (
	"errors"
	goIpfs "github.com/ipfs/go-ipfs-api"
	"io"
	"log"
)

var ErrCannotConnetToNode = errors.New("Check if your node(Ipfs daemon) is running. Cannot connect with node")

func main() {
	ipfs := goIpfs.NewLocalShell()
	if ipfs == nil {
		log.Fatal(ErrCannotConnetToNode)
	}
	// ipfs cat doesn't point to the local filesystem
	path := "bafybeif2ewg3nqa33mjokpxii36jj2ywfqjpy3urdh7v6vqyfjoocvgy3a" // points to a directory
	data, err := ipfs.Refs(path, true)                                    // retrieve the cid in this directory
	if err != nil {
		log.Println(err)
	}

	rc, err := ipfs.Cat(<-data)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(rc)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(bytes))

	err = rc.Close()
	if err != nil {
		log.Println(err)
	}
}
