// @Author: liyongzhen
// @Description:
// @File: func
// @Date: 2022/12/6 22:11

package data

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
	"math/big"
)

// EventHash 根据事件返回对应的hash数据
func EventHash(op string) common.Hash {
	return crypto.Keccak256Hash([]byte(op))
}

// 处理金额
func ChainAmount(amount *big.Int, decimal uint8) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(amount.String())
	amountHuman := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimal))))
	f, _ := amountHuman.Float64()
	return f
}

// ResLog 请求返回数据
type ResLog struct {
	ContractAddress string     `json:"contract_address"`
	FromAddress     string     `json:"from_address"`               // 交易的from地址
	ToAddress       string     `json:"to_address"`                 // 交易的to地址
	Amount          float64    `json:"amount,omitempty"`           // 转账份数/金额
	Decimal         uint8      `json:"-"`                          // erc20的精度
	TokenId         *big.Int   `json:"token_id,omitempty"`         // NFT id
	ContractType    string     `json:"contract_type"`              // 合约类型
	Event           string     `json:"event"`                      // 操作类型
	LogIndex        uint       `json:"log_index"`                  // 当前log发生的顺序
	BlockHash       string     `json:"block_hash"`                 // 区块hash
	TxHash          string     `json:"tx_hash"`                    // 交易hash
	Amounts         []*big.Int `json:"amounts,omitempty"`          // 针对erc1155的批量转账
	TokenIds        []*big.Int `json:"token_ids,omitempty"`        // 针对erc1155的批量转账
	ApprovalForAll  bool       `json:"approval_for_all,omitempty"` // 针对erc1155和erc721的字段
}

// EthRecord eth金额流水
type EthRecord struct {
	FromAddress string  `json:"from_address"` // from用户地址
	ToAddress   string  `json:"to_address"`   // to用户地址
	Amount      float64 `json:"amount"`       // 转账金额
	FeeAmount   float64 `json:"fee_amount"`   // 手续费金额
	FromAmount  float64 `json:"from_amount"`  // 流水之后的金额
	ToAmount    float64 `json:"to_amount"`    // 流水之后的金额
	TxID        string  `json:"tx_id"`        // 公链的交易id
}
