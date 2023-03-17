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
		Provider: "xxx",
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

func TestCS_GetETHLogsByNumber(t *testing.T) {
	// 注意这里需要使用存档节点或者全节点，不然会报错
	c := client.C{
		Provider: "xxx",
	}
	cc, err := ethclient.Dial(c.Provider)
	if err != nil {
		t.Fatal(err)
	}
	cs := CS{
		Client: cc,
		Ctx:    context.Background(),
	}
	blocks, err := cs.GetETHLogsByNumber(16590442)
	if err != nil {
		t.Fatal(err)
	}
	jsonS, _ := json.Marshal(blocks)
	t.Log(string(jsonS))
}

func TestCS_GetErc20TransferBlocks(t *testing.T) {
	c := client.C{
		Provider: "xxxx",
	}
	cc, err := ethclient.Dial(c.Provider)
	if err != nil {
		t.Fatal(err)
	}
	cs := CS{
		Client:       cc,
		FromBlock:    16840054,
		ToBlock:      16840054,
		Ctx:          context.Background(),
		TopicEvent:   TransferEvent,
		ContractType: ContractType20,
	}
	blocks, err := cs.GetBlocks()
	if err != nil {
		t.Fatal(err)
	}
	jsonS, _ := json.Marshal(blocks)
	t.Log(string(jsonS))
}
