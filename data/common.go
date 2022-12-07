// @Author: liyongzhen
// @Description:
// @File: common
// @Date: 2022/12/6 22:04

package data

import "errors"

var (
	paramsErr = errors.New("missing parameter")
)

// 合约类型
const (
	ContractType20   = "ERC20"
	ContractType721  = "ERC721"
	ContractType1155 = "ERC1155"
)

// 合约事件
const (
	DeployContract = "DeployContract" // 部署合约
	TransferSingle = "TransferSingle"
	TransferBatch  = "TransferBatch"
	ApprovalForAll = "ApprovalForAll"
	Transfer       = "Transfer"
	Approval       = "Approval"
)

// DeployContractHash 部署合约的事件hash
const DeployContractHash = "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0"

// ERC1155事件
const (
	ERC1155TransferSingle = "TransferSingle(address,address,address,uint256,uint256)"
	ERC1155TransferBatch  = "TransferBatch(address,address,address,uint256[],uint256[])"
)

// ERC721/20/1155公共事件
const (
	TransferEvent      = "Transfer(address,address,uint256)"
	ApprovalEvent      = "Approval(address,address,uint256)"
	ApproveForAllEvent = "ApprovalForAll(address,address,bool)"
)
