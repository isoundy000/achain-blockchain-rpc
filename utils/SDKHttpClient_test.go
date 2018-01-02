package utils

import (
	"2018go-workspace/achain-blockchain/blockchain"
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	var params []string
	//result, _ := Post("http://172.16.33.16:18888/rpc", "admin:123456", "info", params)
	result, _ := Post("http://172.16.33.16:18888/rpc", "admin:123456", "blockchain_get_block_count", params)
	//result := Post(common.WALLET_RPC, common.WALLET_NAME_PASSWORD, "info", params)
	fmt.Println(result)
	//info := blockchain.Info(result)
	//fmt.Println(info)
	//demo()

}

func demo() {
	a := 1
	fmt.Println(a)

	inChan := make(chan int, 2)

	inChan <- 2

	var b int

	b = <-inChan

	close(inChan)
	fmt.Println(b)

}
