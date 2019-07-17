package main

import(
  "strconv"
  "fmt"
  "github.com/lukzhang/GoLang_Blockchai/blockchain"
)



func main(){
  chain := blockchain.InitBlockChain()

  chain.AddBlock("First Block after Genesis")
  chain.AddBlock("Second Block after Genesis")
  chain.AddBlock("Third Block after Genesis")

  for _, block := range chain.blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data in Block: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)

    pow := blockchain.NewProof(block)
    fmt.Printf("PoW: %s\n", strconv.Formatool(pow.Validate()))
    fmt.Println()
  }
}
