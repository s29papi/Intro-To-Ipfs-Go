package main

import (
	"errors"
	goIpfs "github.com/ipfs/go-ipfs-api"
	"log"
)

var ErrCannotConnetToNode = errors.New("Check if your node(Ipfs daemon) is running. Cannot connect with node")

func main() {
	ipfs := goIpfs.NewLocalShell()
	if ipfs == nil {
		log.Fatal(ErrCannotConnetToNode)
	}
	err := ipfs.Get("bafybeif2ewg3nqa33mjokpxii36jj2ywfqjpy3urdh7v6vqyfjoocvgy3a", "files")
	if err != nil {
		log.Fatal(err)
	}
	info := "Info \r\n Successfully retreived file. \r\n Check in the files directory for this file which we just retreived \r\n N.B You are also now hosting the folder and its contents for others to retrieve."
	log.Println(info)
}
