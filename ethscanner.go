package main

import (
	"bufio"
	"encoding/json"
	_ "ethscanner/memstore" // initialize mem storage
	"ethscanner/parser"
	"fmt"
	"log"
	"os"
)

const URL = "https://cloudflare-eth.com"

// "0xdac17f958d2ee523a2206206994597c13d831ec7" // usdt contract address

func main() {
	p := parser.New(URL)
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		switch in.Text() {
		case "exit":
			return
		case "block":
			block := p.GetCurrentBlock()
			log.Println(fmt.Sprintf("Last processed block: %d", block))
		case "tx":
			log.Println("Enter address:")
			if in.Scan() {
				address := in.Text()
				log.Println(fmt.Sprintf("Transactions for address %s:", address))
				transactions := p.GetTransactions(address)
				if len(transactions) == 0 {
					log.Println("Not found")
				} else {
					marshal, _ := json.Marshal(transactions)
					log.Println(fmt.Sprintf("%s", string(marshal)))
				}
			}
		case "sub":
			log.Println("Enter address:")
			if in.Scan() {
				address := in.Text()
				if p.Subscribe(address) {
					log.Println(fmt.Sprintf("Address %s has been subscribed", address))
				} else {
					log.Println(fmt.Sprintf("Address %s has already been subscribed", address))
				}
			}
		default:
			log.Println("Unknown command")
		}
	}
}
