package main

import (
	"fmt"

	"main.go/Function"
)

// https://www.blockchain.com/pt/api/blockchain_api

func main() {

	var hash string = "b6f6991d03df0e2e04dafffcd6bc418aac66049e2cd74b80f14ac86db1e3f0da"

	fmt.Printf("API Response as struct %+v\n", Function.GetTransaction(hash))

}
