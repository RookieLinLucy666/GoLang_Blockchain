# GoLang_Blockchain
Blockchain made in Go with Proof of Work (PoW) using sha256


git clone https://github.com/lukzhang/GoLang_Blockchain.git


# Open 3 terminals and go into root folder. We are going have 3 nodes. One will create a genesis block, one will mine and check for transactions, and the other will send a token to another wallet

cd golang-blockchain


# In first terminal

set NODE_ID=3000

go run main.go createwallet


# In second terminal

set NODE_ID=4000

go run main.go createwallet


# In third terminal

set NODE_ID=5000

go run main.go createwallet


# In first terminal

go run main.go createblockchain -address <TERMINAL1 WALLET>
  
  # Go into /tmp and copy 'blocks_3000' folder for 4000 and 5000
  
  cd tmp
  
  cp -r blocks_3000 blocks_4000
  
  cp -r blocks_3000 blocks_5000
  
  cp -r blocks_3000 blocks_gen
  
  # Go back to root folder and give second account 10 tokens by mining
  
  cd ..
  
  go run main.go send -from <TERMINAL1 WALLET> -to <TERMINAL2 WALLET> -amount 10 -mine
  
  # start node
  
  go run main.go startnode
  
  
# In second terminal, start node as well

go run main.go

# In third terminal, start miner

go run main.go startnode -miner <TERMINAL3 WALLET>
  
 # In second terminal, ctrl+c to stop the node and send a transaction from the 10 tokens it received earlier
 
 go run main.go send -from <TERMINAL2 WALLET> -to <TERMINAL3 WALLET> -amount 3
  
  # You will see 3 tokens being sent as the miner mines a block. Terminal 1 which is running the node will show the block hash that was mined
  
  # In second terminal, you can sync to the blockchain by running the node again
  
  go run main.go startnode
  
  # You can view the blockchain by printing if you wish
  
  go run main.go printchain
