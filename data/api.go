// @Author: liyongzhen
// @Description:
// @File: api
// @Date: 2022/12/6 21:46

package data

import "C"
import (
	"context"
	"encoding/binary"
	"github.com/chimojiacai/eth-block-data/ABI"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
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
	if err := c.checkParams(); err != nil {
		return nil, err
	}
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
	switch log.Topics[0] {
	case common.HexToHash(DeployContractHash): // 创建合约
		res.FromAddress = common.HexToAddress(log.Topics[2].Hex()).String() // 创建者
		res.ToAddress = common.HexToAddress(log.Topics[1].Hex()).String()   // 0地址
		res.Event = DeployContract
		c.checkContractTypeIs721Or1155(res)
		if res.ContractType == "" {
			return nil
		}
	case EventHash(ERC1155TransferSingle): // erc1155 单次转账
		res.ContractType = ContractType1155
		res.Event = TransferSingle
		doAddress23(log.Topics, res)
		contractAbi, err := abi.JSON(strings.NewReader(ABI.ERC1155ABI)) // 合约数据
		if err != nil {
			return nil
		}
		singleLogData := struct {
			Id    *big.Int `json:"id"`
			Value *big.Int `json:"value"`
		}{}
		if err := contractAbi.UnpackIntoInterface(&singleLogData, res.Event, log.Data); err != nil {
			return nil
		}
		res.Amount = float64(singleLogData.Value.Int64())
		res.TokenId = singleLogData.Id
	case EventHash(ERC1155TransferBatch): // erc1155 批量转账
		res.ContractType = ContractType1155
		res.Event = TransferBatch
		doAddress23(log.Topics, res)
		contractAbi, err := abi.JSON(strings.NewReader(ABI.ERC1155ABI)) // 合约数据
		if err != nil {
			return nil
		}
		logData := struct {
			Ids    []*big.Int `json:"id"`
			Values []*big.Int `json:"value"`
		}{}
		if err := contractAbi.UnpackIntoInterface(&logData, res.Event, log.Data); err != nil {
			return nil
		}
		res.Amounts = logData.Values
		res.TokenIds = logData.Ids
	case EventHash(ApproveForAllEvent): // erc1155/721 批量授权
		c.checkContractTypeIs721Or1155(res)
		if res.ContractType == "" {
			return nil
		}
		res.Event = ApprovalForAll
		doAddress12(log.Topics, res)
		res.ApprovalForAll = common.BytesToHash(log.Data).Big().Cmp(big.NewInt(1)) == 0
	case EventHash(TransferEvent): // erc721/20 转账
		c.checkContractTypeIs721Or20(res)
		if res.ContractType == "" {
			return nil
		}
		res.Event = Transfer
		if res.ContractType == ContractType20 {
			doAddress12(log.Topics, res)
			dValue := common.BytesToHash(log.Data)
			res.Amount = ChainAmount(dValue.Big(), res.Decimal) // 真实金额
		} else if res.ContractType == ContractType721 {
			doAddress12(log.Topics, res)
			res.TokenId = log.Topics[3].Big()
		}
	case EventHash(ApprovalEvent): // erc721/20 单次授权
		c.checkContractTypeIs721Or20(res)
		if res.ContractType == "" {
			return nil
		}
		res.Event = Approval
		if res.ContractType == ContractType721 {
			doAddress12(log.Topics, res)
			res.TokenId = log.Topics[3].Big()
		} else if res.ContractType == ContractType20 {
			doAddress12(log.Topics, res)
			dValue := common.BytesToHash(log.Data)
			res.Amount = ChainAmount(dValue.Big(), res.Decimal) // 真实金额
		}
	default:
		return nil
	}
	return res
}

// 判断合约类型
func (c *CS) checkContractTypeIs721Or1155(res *ResLog) {
	// 判断erc721
	e721, err := ABI.NewABIErc721Caller(common.HexToAddress(res.ContractAddress), c.Client) // 这里是合约地址
	if err == nil {
		var blob [4]byte
		binary.BigEndian.PutUint32(blob[:], 0x80ac58cd)
		supportsInterface, err := e721.SupportsInterface(nil, blob)
		if err == nil && supportsInterface {
			res.ContractType = ContractType721
			return
		}
	}
	// 判断erc1155
	e1155, err := ABI.NewABIErc1155Caller(common.HexToAddress(res.ContractAddress), c.Client) // 这里是合约地址
	if err == nil {
		var blob [4]byte
		binary.BigEndian.PutUint32(blob[:], 0xd9b67a26) // 0xd9b67a26=1155
		supportsInterface, err := e1155.SupportsInterface(nil, blob)
		if err == nil && supportsInterface {
			res.ContractType = ContractType1155
			return
		}
	}
}

// 判断合约类型
func (c *CS) checkContractTypeIs721Or20(res *ResLog) {
	// 判断erc721
	e721, err := ABI.NewABIErc721Caller(common.HexToAddress(res.ContractAddress), c.Client) // 这里是合约地址
	if err == nil {
		var blob [4]byte
		binary.BigEndian.PutUint32(blob[:], 0x80ac58cd)
		supportsInterface, err := e721.SupportsInterface(nil, blob)
		if err == nil && supportsInterface {
			res.ContractType = ContractType721
			return
		}
	}

	// 判断erc20
	if c.checkContractIsErc20(res) {
		res.ContractType = ContractType20
		return
	}
}

// 判断合约地址是否为erc20
// 判断当前合约是否实现了：
// totalSupply()
// balanceOf(address)
// transfer(address,uint256)
// transferFrom(address,address,uint256)
// approve(address,uint256)
// allowance(address,address)
func (c *CS) checkContractIsErc20(res *ResLog) bool {
	erc20Caller, err := ABI.NewABIErc20Caller(common.HexToAddress(res.ContractAddress), c.Client)
	if err != nil {
		return false
	}
	d, err := erc20Caller.Decimals(nil)
	if err != nil {
		return false
	}
	if d == 0 {
		return false
	}
	res.Decimal = d
	_, err = erc20Caller.TotalSupply(nil)
	if err != nil {
		return false
	}
	_, err = erc20Caller.Symbol(nil)
	if err != nil {
		return false
	}
	name, err := erc20Caller.Name(nil)
	if err != nil || name == "" {
		return false
	}
	//_, err = erc20Caller.BalanceOf(nil, common.HexToAddress(res.ContractAddress))
	//if err != nil {
	//	return false
	//}
	return true
}

// 处理地址
func doAddress12(logs []common.Hash, res *ResLog) {
	res.FromAddress = common.HexToAddress(logs[1].Hex()).String()
	res.ToAddress = common.HexToAddress(logs[2].Hex()).String()
}

// 处理地址
func doAddress23(logs []common.Hash, res *ResLog) {
	res.FromAddress = common.HexToAddress(logs[2].Hex()).String()
	res.ToAddress = common.HexToAddress(logs[3].Hex()).String()
}

// GetETHLogsByNumber 通过高度获取链上ETH流水记录
// 可在自己的业务逻辑中过滤得到自己想要的address
func (c *CS) GetETHLogsByNumber(number int64) ([]*EthRecord, error) {
	logs, err := c.Client.BlockByNumber(c.Ctx, big.NewInt(number))
	if err != nil {
		return nil, err
	}
	if logs == nil || len(logs.Transactions()) == 0 {
		return nil, nil
	}
	var res []*EthRecord
	for _, transaction := range logs.Transactions() {
		tmp := &EthRecord{}
		tmp.Amount = ChainAmount(transaction.Value(), 18) // 转账金额
		msg, err := transaction.AsMessage(types.NewEIP155Signer(transaction.ChainId()), nil)
		if err != nil {
			return nil, err
		}
		tmp.FromAddress = msg.From().Hex() // from地址
		if transaction.To() != nil {
			tmp.ToAddress = transaction.To().String() // to地址
		}
		tmp.TxID = transaction.Hash().Hex() // 交易hash
		// 获取交易详情
		transactionReceipt, err := c.Client.TransactionReceipt(c.Ctx, transaction.Hash())
		if err != nil {
			return nil, err
		}
		// 真实gas费用
		tmp.FeeAmount = ChainAmount(big.NewInt(int64(transactionReceipt.GasUsed)), 9)
		// 获取from和to的地址的当时区块高度的金额
		atFrom, err := c.Client.BalanceAt(c.Ctx, msg.From(), transactionReceipt.BlockNumber)
		if err != nil {
			return nil, err
		}
		tmp.FromAmount = ChainAmount(atFrom, 18)
		if transaction.To() != nil {
			atTo, err := c.Client.BalanceAt(c.Ctx, *transaction.To(), transactionReceipt.BlockNumber)
			if err != nil {
				return nil, err
			}
			tmp.ToAmount = ChainAmount(atTo, 18)
		}

		res = append(res, tmp)
	}
	return res, nil
}
