# ExSat Bridge EOS 集成示例 (TypeScript)

这是一个使用TypeScript集成EOS区块链的简单示例，主要演示如何调用EOS合约和查询合约表数据。

## 功能

- 使用环境变量配置EOS账号和私钥
- 通过HTTP API调用EOS合约的actions
- 通过HTTP API查询EOS合约表数据

## 安装

```bash
# 安装依赖
npm install

# 编译TypeScript
npm run build
```

## 配置

在项目根目录创建一个`.env`文件，设置以下环境变量：

```
EOS_ACCOUNT=youreosaccount
EOS_PRIVATE_KEY=yourprivatekey
EOS_NODE_URL=https://eos.greymass.com
EOS_CONTRACT=yourcontract
PORT=3000
```

## 运行

```bash
npm start
```

## API接口

### 1. 调用合约Action

**请求**:
```
POST /api/contract/action
Content-Type: application/json

{
  "action": "transfer",
  "data": {
    "from": "youreosaccount",
    "to": "receiver",
    "quantity": "1.0000 EOS",
    "memo": "test transfer"
  }
}
```

**响应**:
```json
{
  "success": true,
  "transactionId": "1a2b3c...",
  "blockNum": 123456789,
  "data": {
    // 完整的交易结果
  }
}
```

### 2. 查询合约表

**请求**:
```
GET /api/contract/table/accounts?scope=youreosaccount&limit=10
```

**响应**:
```json
{
  "success": true,
  "rows": [
    {
      "balance": "10.0000 EOS"
    }
  ],
  "more": false
}
```

## 使用示例

### 调用合约示例 (使用curl)

```bash
curl -X POST http://localhost:3000/api/contract/action \
  -H "Content-Type: application/json" \
  -d '{
    "action": "transfer",
    "data": {
      "from": "youreosaccount",
      "to": "receiver",
      "quantity": "1.0000 EOS",
      "memo": "test transfer"
    }
  }'
```

### 查询表示例 (使用curl)

```bash
curl -X GET "http://localhost:3000/api/contract/table/accounts?scope=youreosaccount"
```
