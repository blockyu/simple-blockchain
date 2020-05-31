package main

import (
	"simple-blockchain/core"
)

func main() {
	bc := core.NewBlockchain()
	defer bc.DB.Close()

	cli := CLI{bc}
	cli.Run()
}
