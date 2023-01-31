package main

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
)

var (
    ctx         = context.Background()
    url         = "https://mainnet.infura.io/v3/aada9accef8a41fd8f503b051c6ed8bc"
    client, _ = ethclient.DialContext(ctx, url)
)

func currentBlock() {
    block, err := client.BlockByNumber(ctx, nil)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(block.Number())
}

func main() {
	currentBlock()
}