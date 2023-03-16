package main

import (
	"fmt"
	. "github.com/wslzwps/go-web3/rpc"
	"log"
	"time"
)

func main() {
	client, err := NewClient("wss://ws-nd-371-746-352.p2pify.com/35176065ca87c60c023e583f8dca14a8", "")
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Subscribe("newPendingTransactions", func(data []byte) {
		fmt.Printf("data %s\n", data)
	})
	if err != nil {
		log.Fatal(err)
	}
	<-time.After(time.Minute)
}
