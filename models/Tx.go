package models

import "time"

type Transaction struct {
	ID        string    `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Amount    int       `json:"amount"`
	Fee       float64   `json:"fee"`
	Timestamp time.Time `json:"timestamp"`
}

type TransactionCheckout struct {
	TxIDs [] string `json:"txID"`
	MerkleRoot string `json:"merkle_root"`
	Miner string `json:"miner"`
	Timestamp time.Time `json:"timestamp"`
	IsGenesis bool `json:"is_genesis"`
	Nonce int `json:"nonce"`
}