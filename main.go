package main

import (
	bc "block-chain/block"
	"bufio"
	"fmt"
	"os"
)

func help() {
	fmt.Println("Type c for add a newblock")
	fmt.Println("Type p for print all the blockChain")
	fmt.Println("Type q for exiting")
}

func main() {
	fmt.Println("BlockChain Project Test")
	fmt.Println("Type h for help")

	NewBlockChain := bc.CreateBlockChain()

	var op string

	for true {
		fmt.Scanln(&op)
		if op == "h" {
			help()
		} else if op == "c" {
			fmt.Println("Entering your Data")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()
			NewBlockChain.AddBlock(data)
			fmt.Println("Success")
		} else if op == "p" {
			for _, block := range NewBlockChain.Blocks {
				fmt.Printf("Prev Hash: %x \n", block.PrevBlockHash)
				fmt.Printf("Data: %s \n", block.Data)
				fmt.Printf("Hash: %x \n", block.Hash)
				fmt.Printf("Pow: %t \n", bc.NewProofOfWork(block).Validate())
				fmt.Println()
			}

		} else if op == "q" {
			break
		} else {
			fmt.Println("Please Enter your option")
		}
	}
}
