package main

import (
	"flag"
	"fmt"
	"os"
	"simple-blockchain/core"

)

type  CLI struct {
	bc *core.Blockchain
}

func (cli *CLI) validateArgs() {

}
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err!=nil {
			fmt.Printf("%s", err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err!=nil {
			fmt.Printf("%s", err)
		}
	default:
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("success")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
		fmt.Printf("CurHash: %x\n", block.Hash)
		fmt.Printf("Block: %s\n", block.Data)
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			break
		}

	}
}