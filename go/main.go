package main

import (
	"fmt"

	"main.go/Function"
)

// https://www.blockchain.com/pt/api/blockchain_api

func main() {

	var hash string = "0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103"

	fmt.Printf("API Response as struct %+v\n", Function.GetBloco(hash))

}
