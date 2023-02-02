package main

import (
    "context"
    "github.com/ethereum/go-ethereum/ethclient"
    "log"
	"github.com/ethereum/go-ethereum/common"
    "bufio"
	"fmt"
	"os"
    "regexp"
    "strings"
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/aada9accef8a41fd8f503b051c6ed8bc")
	if err != nil {
		fmt.Println("Error connecting to Ethereum node:", err)
		return
	}
    re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
    
    fmt.Print("Enter your wallet address: ")
	reader := bufio.NewReader(os.Stdin)
	inputData, _ := reader.ReadString('\n')
    address := common.HexToAddress(strings.TrimSpace(inputData))

    match := re.MatchString(address.String())

    addressT := common.HexToAddress(address.String())
	if addressT == (common.Address{}) || !match {
		fmt.Println("Invalid address:", inputData)
        os.Exit(1)
	}

    // Retrieve the code associated with the address
	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Check the length of the code
	if len(code) > 0 {
        fmt.Println("Address is a contract address", address)
    } else {
        fmt.Println("Address is an EOA", address)
	}
}