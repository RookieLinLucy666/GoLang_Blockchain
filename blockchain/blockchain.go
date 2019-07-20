package blockchain

import (
  "fmt"

  "github.com/dgraph-io/badger"
)

const(
  dbPath = "./tmp/blocks"
)

type BlockChain struct{
  // Blocks []*Block
  LastHash []byte
  Database *badger.DB
}

type BlockChainIterator struct {
  CurrentHash []byte
  Database *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string){
  // prevBlock:= chain.Blocks[len(chain.Blocks)-1]
  // new := CreateBlock(data, prevBlock.Hash)
  // chain.Blocks = append(chain.Blocks, new)

  var lastHash[]byte

  err := chain.Database.View(func(txn *badger.Txn) error{
    item, err := txn.Get([]byte("lh"))
    Handle(err)
    lastHash, err = item.Value()

    return err
  })
  Handle(err)

  newBlock := CreateBlock(data, lastHash)

  err = chain.Database.Update(func(txn *badger.Txn) error{
    err := txn.Set(newBlock.Hash, newBlock.Serialize())
    Handle(err)
    err = txn.Set([]byte("lh"), newBlock.Hash)

    chain.LastHash = newBlock.Hash

    return err
  })
  Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator{
  iter:=&BlockChainIterator{chain.LastHash, chain.Database}

  return iter
}

func (iter *BlockChainIterator) Next() *Block{
  var block *Block
  err := iter.Database.View(func(txn *badger.Txn) error{
    item, err := txn.Get(iter.CurrentHash)
    encodedBlock, err := item.Value()
    block = Deserialize(encodedBlock)

    return err
  })
  Handle(err)

  iter.CurrentHash = block.PrevHash

  return block
}

// type Block struct {
//   Hash []byte
//   Data []byte
//   PrevHash []byte
//   Nonce int
// }
//
// func CreateBlock(data string, prevHash []byte) *Block{
//   block := &Block{[]byte{}, []byte(data), prevHash, 0}
//   pow := NewProof(block)
//   nonce, hash := pow.Run()
//
//   block.Hash = hash[:]
//   block.Nonce = nonce
//
//   return block
// }
//
// func Genesis() *Block{
//   return CreateBlock("Genesis", []byte{})
// }
