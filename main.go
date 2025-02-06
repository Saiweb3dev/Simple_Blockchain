package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"

	"Simple_Blockchain/routes"
	"Simple_Blockchain/handlers"
)


func main(){

	var BlockChain *handlers.Blockchain

	BlockChain = handlers.NewBlockchain()

	r := routes.Router(BlockChain)

	go func() {
		for _, block := range BlockChain.Blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block, "", "  ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
	}
}()
 
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}