// @Author: liyongzhen
// @Description:
// @File: func
// @Date: 2022/12/6 22:11

package data

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// EventHash 根据事件返回对应的hash数据
func EventHash(op string) common.Hash {
	return crypto.Keccak256Hash([]byte(op))
}

// ResLog 请求返回数据
type ResLog struct {
	ContractAddress string   `json:"contract_address"`
	FromAddress     string   `json:"from_address"`       // 交易的from地址
	ToAddress       string   `json:"to_address"`         // 交易的to地址
	Amount          float64  `json:"amount"`             // 转账份数/金额
	TokenId         *big.Int `json:"token_id,omitempty"` // NFT id
	ContractType    string   `json:"contract_type"`      // 合约类型
	OperationType   string   `json:"operation_type"`     // 操作类型
	LogIndex        uint     `json:"log_index"`          // 当前log发生的顺序
	BlockHash       string   `json:"block_hash"`         // 区块hash
	TxHash          string   `json:"tx_hash"`            // 交易hash
}
