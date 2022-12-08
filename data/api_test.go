// @Author: liyongzhen
// @Description:
// @File: api_test
// @Date: 2022/12/7 11:05

package data

import (
	"context"
	"encoding/json"
	"github.com/chimojiacai/eth-block-data/client"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestCS_GetBlocks(t *testing.T) {
	c := client.C{
		Provider: "https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7",
	}
	cc, err := ethclient.Dial(c.Provider)
	if err != nil {
		t.Fatal(err)
	}
	cs := CS{
		Client:    cc,
		FromBlock: 16137291,
		ToBlock:   16137291,
		Ctx:       context.Background(),
	}
	blocks, err := cs.GetBlocks()
	if err != nil {
		t.Fatal(err)
	}
	jsonS, _ := json.Marshal(blocks)
	t.Log(string(jsonS))
}
