# exSat Bridge Go Implementation Example

This directory contains an example implementation for integrating with the exSat Bridge using Go. The example demonstrates how to request a Bitcoin deposit address and query Bitcoin deposit address mappings.

## Features

- Create Bitcoin deposit address mappings through exSat Bridge
- Query Bitcoin deposit addresses for EVM recipient addresses
- Provide RESTful API endpoints for easy application integration

## Prerequisites

- Go 1.16 or higher
- Registered exSat Bridge developer account and permission ID
- Access to EOS network nodes

## Installation

1. Clone the repository:

```bash
git clone https://github.com/exsat-network/exSat-bridge-integration-example.git
cd exSat-bridge-integration-example/golang
```

2. Install dependencies:

```bash
go mod download
```

## Configuration

Before running the example, you need to configure environment variables or create a configuration file. The example supports configuration through environment variables or a `.env` file:

```
EOS_ACCOUNT=yourAccount
EOS_PRIVATE_KEY=yourPrivateKey
EOS_NODE_URL=https://your-eos-node-url
BRDGMNG_PERMISSION_ID=0
BTC_BRIDGE_CONTRACT=brdgmng.xsat
PORT=3000
RESOURCE_PAYMENT=false
```

### Configuration Parameters:

- `EOS_ACCOUNT`: Your EOS account name
- `EOS_PRIVATE_KEY`: Your EOS private key
- `EOS_NODE_URL`: URL address of the EOS node
- `BRDGMNG_PERMISSION_ID`: exSat BTC Bridge permission id
- `BTC_BRIDGE_CONTRACT`: exSat Bridge contract account name
- `PORT`: Port for the API service to listen on
- `RESOURCE_PAYMENT`: Whether to use resource payment mode (boolean)

## Running

Compile and run the example service:

```bash
go run main.go
```

## API Endpoints

After starting the service, the following API endpoints are available:

### Health Check

```
GET /api/health
```

Response example:
```json
{
  "status": "ok",
  "timestamp": "2023-08-01T12:34:56Z"
}
```

### Apply for Bitcoin Deposit Address

```
POST /api/brdgmng/appaddrmap
Content-Type: application/json

{
  "recipient_address": "0x123456789abcdef123456789abcdef123456789a",
  "remark": "Test address"
}
```

Response example:
```json
{
  "success": true,
  "transaction": {
    "transaction_id": "1234567890abcdef1234567890abcdef"
    // ...other transaction details
  }
}
```

### Query Bitcoin Deposit Address

```
GET /api/brdgmng/deposit-address/{recipientAddress}
```

Response example:
```json
{
  "recipientAddress": "0x123456789abcdef123456789abcdef123456789a",
  "depositAddress": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
}
```

## Code Structure

- `main.go` - Application entry point and service initialization
- `config/` - Configuration files and environment variable handling
- `services/` - Core service implementations
  - `api_service.go` - REST API service implementation
  - `eos_service.go` - EOS blockchain interaction service
- `utils/` - Utility functions and constants

## Error Handling

The example code includes basic error handling, including:
- API request parameter validation
- Contract call error handling
- Blockchain interaction exception handling

## Production Considerations

This example is for demonstration purposes only. When using in a production environment, consider:
1. Implementing appropriate security measures (e.g., API keys, rate limiting)
2. Adding comprehensive logging and monitoring
3. Implementing retry mechanisms to handle temporary blockchain failures
4. Consider using message queues for asynchronous transaction processing

## Support

If you have any questions, please contact the exSat team:
- Developer support email: support@exsat.org