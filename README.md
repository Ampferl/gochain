# Blockchain written in GO
`gochain` is a simple proof-of-work blockchain written in go.  
## How to run the gochain network
First, set `NODE_ID` to 3000 (`export NODE_ID=3000`) in the first terminal window. 
Create a wallet and a new blockchain on the NODE 3000 terminal:

```shell
$ go run .createblockchain -address <address3000>
```  

After that, the blockchain will contain only the genesis block. We need to save the block and use it in other nodes.  

```shell
$ cp blockchain_3000.db blockchain_genesis.db 
```
---
Next, open a the terminal window of your second Node and set node ID to 3001. This will be a wallet node. Generate some addresses with `go run .createwallet`.

Back to the first window send some coins to the wallet addresses:

```shell
$ go run . send -from <address3000> -to <address3001> -amount 10 -mine
```
`-mine` flag means that the block will be immediately mined by the same node. 

Start the node:

```shell
$ go run . startnode
```
The node must be running until the end of the scenario.

 
Start the blockchain of NODE 3001 with the genesis block saved above:

```shell
$ cp blockchain_genesis.db blockchain_3001.db
```
Run the node:

```shell
$ go run . startnode
```
It'll download all the blocks from the 3000 node. To check that everything's ok, stop the node and check the balances:

```shell
$ go run . getbalance -address <address3001>
Balance of '<address3001>': 10
```
Also, you can check the balance of the NODE 3000 address, because the node 3001 now has its blockchain:

```shell
$ go run . getbalance -address <address3000>
Balance of '<address3000>': 10
```

--- 

Open a new terminal window and set its ID to 3002, and generate a wallet. This will be a miner node. Initialize the blockchain:

```shell
$ cp blockchain_genesis.db blockchain_3002.db
```
And start the node:

```shell
$ go run . startnode -miner <address3002>
```

Send some coins from NODE 3001:

```shell
$ go run . send -from <address3001> -to <address3002> -amount 1
$ go run . send -from <address3001> -to <address3000> -amount 1
```


Quickly! Switch to the miner node 3002 and see it mining a new block! Also, check the output of the central node.


Switch to the node 3001 and start it:

```shell
$ go run . startnode
```
It'll download the newly mined block!

Stop it and check balances:

```shell
$ go run . getbalance -address <address3001>
Balance of 'WALLET_1': 8
```

## Ressources
- [Blockchain für Entwickler](https://www.amazon.de/Blockchain-f%C3%BCr-Entwickler-Programmierung-Praxisbeispielen/dp/3836263904)