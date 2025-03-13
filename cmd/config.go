package cmd

type Config struct {
	RpcURL                string `json:"rpcURL"`
	FromPrivateKey        string `json:"fromPrivateKey"`
	ToAddress             string `json:"toAddress"`
	SendValue             uint64 `json:"sendValue"`
	GasPrice              uint64 `json:"gasPrice"`
	GasLimit              uint64 `json:"gasLimit"`
	TxAmount              int    `json:"txAmount"`
	TxSendIntervalSeconds int    `json:"txSendIntervalSeconds"`
	WaitForReceipt        bool   `json:"waitForReceipt"`
	Type                  int    `json:"type"`
}
