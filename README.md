# Tx Type Sender

Send Type Tx's

## Usage

- Step 1: Edit `sender-cfg.json` (change type to either 1/2/3/4)
- Step 2: `go run main.go`

## Configuration

```json
{
  "rpcURL": "http://127.0.0.1:8123",
  "fromPrivateKey": "yourprivkey",
  "toAddress": "0x",
  "sendValue": 10000,
  "gasPrice": 1000000000,
  "gasLimit": 30000000,
  "txAmount": 10,
  "txSendIntervalSeconds": 0,
  "waitForReceipt": false,
  "Type": 1
}
```