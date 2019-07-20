# GoLang_Blockchain
Blockchain made in Go with Proof of Work (PoW) using sha256

in root folder

Make sure ./tmp/blocks is empty so you can start creating a new chain


#First time running blockchain, create genesis block


go run main.go

#View the blocks so far on chain


go run main.go print

#add a new block with a message


go run main.go add -block "I added a block!"

"View your change


go run main.go print
