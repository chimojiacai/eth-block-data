# eth-block-data
Analyze block data
## 分析eth区块数据
根据区块高度扫描分析以太坊链（私有链、公链、测试链、侧链等）合约部署及交易信息
目前只实现了（`ERC20`，`ERC721`，`ERC1155`）三种协议的部署和交易信息。
主要使用`eth_getLogs`来获取log数据
## 测试案例
```azure
func TestCS_GetBlocks(t *testing.T) {
    c := client.C{
        Provider: "http://xxxxx", // 节点链接
    }
    cc, err := ethclient.Dial(c.Provider)
    if err != nil {
        t.Fatal(err)
    }
    cs := CS{
        Client:    cc,
        FromBlock: 1519251,
        ToBlock:   1519251,
        Ctx:       context.Background(),
    }
    blocks, err := cs.GetBlocks()
    if err != nil {
        t.Fatal(err)
    }
    jsonS, _ := json.Marshal(blocks)
    t.Log(string(jsonS))
}

```

### 返回数据

```azure
[
    {
        "contract_address": "0xEb433F1f5272bBaCCa3BF3f9d9527FbF5010aDe1",
        "from_address": "",
        "to_address": "",
        "amount": 0,
        "contract_type": "ERC20",
        "operation_type": "Approval",
        "log_index": 0,
        "block_hash": "0x0009762f36be20e99fbdaf33679acf653699e6b41572613c6dfea3dab9c280d7",
        "tx_hash": "0xe639325ed420c8ff8a0095735e83e566febbb2ff8b3186ef41e311e3237c0759"
    },
    {
        "contract_address": "0xEb433F1f5272bBaCCa3BF3f9d9527FbF5010aDe1",
        "from_address": "",
        "to_address": "",
        "amount": 0,
        "contract_type": "ERC20",
        "operation_type": "Transfer",
        "log_index": 1,
        "block_hash": "0x0009762f36be20e99fbdaf33679acf653699e6b41572613c6dfea3dab9c280d7",
        "tx_hash": "0xe639325ed420c8ff8a0095735e83e566febbb2ff8b3186ef41e311e3237c0759"
    }
]
```
