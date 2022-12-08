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
		Provider: "xxxx",
		ChainId:  9215,
	}
	cc, err := ethclient.Dial(c.Provider)
	if err != nil {
		t.Fatal(err)
	}
	cs := CS{
		Client:    cc,
		FromBlock: 1521801,
		ToBlock:   1522008,
		Ctx:       context.Background(),
	}
	blocks, err := cs.GetBlocks()
	if err != nil {
		t.Fatal(err)
	}
	jsonS, _ := json.Marshal(blocks)
	t.Log(string(jsonS))
}
