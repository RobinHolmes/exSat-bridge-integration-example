# ExSat Bridge Integration Example (TypeScript)

This is a TypeScript example demonstrating how to integrate with the ExSat Bridge on the EOS blockchain. It shows how to call contract actions and query contract tables.

## Features

- ExSat Bridge API integration
- EOS contract interactions
- BTC deposit address management
- Environment-based configuration

## Prerequisites

- Node.js
- yarn
- EOS account with active permission

## Installation

```bash
# Install dependencies
yarn install

# Build TypeScript
yarn build
```

## Configuration

Create a `.env` file in the project root with the following environment variables:

```
EOS_ACCOUNT=your_eos_account
EOS_PRIVATE_KEY=your_private_key
EOS_NODE_URL=https://rpc-sg.exsat.network
BRDGMNG_PERMISSION_ID=0
RESOURCE_PAYMENT=false
PORT=3000
```

See the `.env.example` file for more details on each configuration parameter.

## Running the Application

```bash
# Development mode
yarn dev

# Production mode
yarn build
yarn start
```

## API Endpoints

### Health Check
```
GET /api/health
```

### Apply for BTC Deposit Address
```
POST /api/brdgmng/appaddrmap
Content-Type: application/json

{
  "recipient_address": "0xYourEvmAddress",
  "remark": "optional note"
}
```

### Get BTC Deposit Address
```
GET /api/brdgmng/deposit-address/:recipientAddress
```

## Example Usage

### Apply for BTC Address (using curl)

```bash
curl -X POST http://localhost:3000/api/brdgmng/appaddrmap \
  -H "Content-Type: application/json" \
  -d '{
    "recipient_address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
    "remark": "test address"
  }'
```

### Get BTC Address (using curl)

```bash
curl -X GET "http://localhost:3000/api/brdgmng/deposit-address/0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
```

## Security Notes

- Never commit your `.env` file containing your private key
- Use proper permission management in production
- Consider using a dedicated API key system for production deployments
