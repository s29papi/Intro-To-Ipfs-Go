package main

import (
	"bytes"
	"errors"
	"fmt"
	goIpfs "github.com/ipfs/go-ipfs-api"
	"log"
	"os"
)

var ErrCannotConnetToNode = errors.New("Check if your node(Ipfs daemon) is running. Cannot connect with node")

func main() {
	ipfs := goIpfs.NewLocalShell()
	if ipfs == nil {
		log.Fatal(ErrCannotConnetToNode)
	}
	fileBytes, err := os.ReadFile("files/cute-puppy.jpg")
	if err != nil {
		log.Fatal(err)
	}
	r := bytes.NewReader(fileBytes)
	cid, err := ipfs.Add(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s", cid)
}

// Info
// puppy image: https://www.google.com/search?q=cute+puppy&sxsrf=ALiCzsYOT3J1pwHy5hMTsVUt8IrsHXC28A:1656939511161&tbm=isch&source=iu&ictx=1&vet=1&fir=Xlf4qHwbPczmKM%252C-_iTcnLa7aD28M%252C_%253BuRKK5LTasSA3DM%252CIel7gHGXEs84MM%252C_%253BcryCV0b0fKPLvM%252C-_iTcnLa7aD28M%252C_%253B1ryxMTgWs8vtUM%252Cb8nd5ak1wbDbqM%252C_%253BqsrJpHzAJBU6uM%252C2ssGwWYAAf269M%252C_%253BpY8Ro_GBr9vrdM%252C2ssGwWYAAf269M%252C_%253BesgXc8A4vAoQiM%252CFUSfTDut5UdV4M%252C_%253BtbCI3AxdYnBSQM%252Cymb7EZRy1fxx-M%252C_%253B8wVms8lnQzUe5M%252C_5QXuchd77nOyM%252C_%253Bk9xReEDKDkeNKM%252C2r6Arj4-hBjhNM%252C_%253Bfob80O8cJcZJjM%252CqNa_o8Sa7Us_pM%252C_%253B2CQDxS6-wVk1BM%252CRcosvHyH6I78fM%252C_%253BISicQ0MBlk1pJM%252CXnNq2DC9yJjNjM%252C_%253BcELtgslghZ_QPM%252C2r6Arj4-hBjhNM%252C_%253BcilC3nARdX40mM%252C5VtXvr9pQCNMrM%252C_&usg=AI4_-kTOYIE2iSgwTmPI7ktXuBVPL8Cw7g&sa=X&ved=2ahUKEwjFj-jXpN_4AhXChc4BHWPHBigQ9QF6BAgREAE#imgrc=uRKK5LTasSA3DM
// By default the go-ipfs add command pins the file upon calling "add"
