# Blockchain written in GO
`gochain` is a simple proof-of-work blockchain written in go.  
## CLI Commands
Create Blockchain:
```shell
$ go run . createblockchain -address <address>
```
Get Balance of a Address:
```shell
$ go run . getbalance -address <address>
```
Print gochain:
```shell
$ go run . printchain
```
Send coins:
```shell
$ go run . send -from <address> -to <address> -amount <amount>
```