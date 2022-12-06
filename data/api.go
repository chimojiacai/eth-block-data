// @Author: liyongzhen
// @Description:
// @File: api
// @Date: 2022/12/6 21:46

package data

import "C"
import (
	"context"
	"ebd/ABI"
	"encoding/binary"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type CS struct {
	Client    *ethclient.Client
	FromBlock int64    // beginning of the queried range, nil means genesis block
	ToBlock   int64    // end of the range, nil means latest block
	Addresses []string // 筛选合约地址
	Ctx       context.Context
}

func (c *CS) checkParams() error {
	if c.FromBlock == 0 && c.ToBlock == 0 && len(c.Addresses) == 0 {
		return paramsErr
	}
	return nil
}

// GetBlocks 获取区块信息
func (c *CS) GetBlocks() ([]*ResLog, error) {
	var as []common.Address
	if len(c.Addresses) > 0 {
		for _, address := range c.Addresses {
			as = append(as, common.HexToAddress(address))
		}
	}
	// 拼接参数读取数据
	query := ethereum.FilterQuery{
		Addresses: as,
		FromBlock: big.NewInt(c.FromBlock),
		ToBlock:   big.NewInt(c.ToBlock),
	}

	logs, err := c.Client.FilterLogs(c.Ctx, query)
	if err != nil {
		return nil, err
	}
	if len(logs) == 0 {
		return nil, nil
	}
	var list []*ResLog
	// 开始分析数据
	for _, v := range logs {
		if len(v.Topics) < 3 {
			continue
		}
		log := c.doLog(v)
		if log == nil {
			continue
		}
		list = append(list, log)
	}

	return list, nil
}

// 处理log数据
func (c *CS) doLog(log types.Log) *ResLog {
	res := &ResLog{
		ContractAddress: log.Address.Hex(),
		LogIndex:        log.Index,
		BlockHash:       log.BlockHash.String(),
		TxHash:          log.TxHash.String(),
	}
	c.checkContractType(res)
	if res.ContractType == "" {
		return nil
	}
	switch log.Topics[0] {
	case common.HexToHash(DeployContractHash): // 创建合约
		res.FromAddress = common.HexToAddress(log.Topics[2].Hex()).String() // 创建者
		res.ToAddress = common.HexToAddress(log.Topics[1].Hex()).String()   // 0地址
		res.OperationType = DeployContract
		// 判断合约类型
	case EventHash(ERC1155TransferSingle): // erc1155 单次转账

	}
	return nil
}

// 判断合约类型
func (c *CS) checkContractType(res *ResLog) {
	// todo 判断erc20
	// 判断erc721
	e721, err := ABI.NewABIErc721Caller(common.HexToAddress(res.ContractAddress), c.Client) // 这里是合约地址
	if err == nil {
		var blob [4]byte
		binary.BigEndian.PutUint32(blob[:], 0x80ac58cd)
		supportsInterface, err := e721.SupportsInterface(nil, blob)
		if err == nil && supportsInterface {
			res.ContractType = ContractType721
		}
		return
	}
	// 判断erc1155
	e1155, err := ABI.NewABIErc1155Caller(common.HexToAddress(res.ContractAddress), c.Client) // 这里是合约地址
	if err == nil {
		var blob [4]byte
		binary.BigEndian.PutUint32(blob[:], 0xd9b67a26)
		supportsInterface, err := e1155.SupportsInterface(nil, blob)
		if err == nil && supportsInterface {
			res.ContractType = ContractType1155
		}
		return
	}
}
