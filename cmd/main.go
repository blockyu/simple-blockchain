package main

import "simple-blockchain/core"

func main() {
	bc := core.NewBlockchain()

	bc.AddBlock("Send 1 BTC To Alice")
	bc.AddBlock("Send 2 BTC To Bob")

	bc.SearchBlock(0, 100)
}
