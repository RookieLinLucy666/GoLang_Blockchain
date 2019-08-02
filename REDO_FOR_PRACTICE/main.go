package main

import (
	"os"

	"github.com/lukzhang/golang-blockchain-redo/wallet"
)

func main() {
	defer os.Exit(0)
	// cmd := cli.CommandLine{}
	// cmd.run()

	w := wallet.MakeWallet()
	w.Address()
}
