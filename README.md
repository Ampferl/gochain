# Blockchain written in GO
`gochain` is a simple proof-of-work blockchain written in go.  
## CLI Commands
Create Address:
```shell
$ go run . createaddress
```
List all Addresses:
```shell
$ go run . listaddresses
```
Create Blockchain with created Address:
```shell
$ go run . createblockchain -address <address>
```
Send coins:
```shell
$ go run . send -from <address> -to <address> -amount <amount>
```
Reindex UTXO:
```shell
$ go run . reindexutxo
```
Get Balance of a Address:
```shell
$ go run . getbalance -address <address>
```
Print gochain:
```shell
$ go run . printchain
```