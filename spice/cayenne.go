package blockchain

//************************************************************
//****GAMEPLAN: Lets see if we can add some state tree mixed with
//some info for now
//************************************************************

import "crypto/sha256"

//???Should we use Merkle-Patricia??
type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}
