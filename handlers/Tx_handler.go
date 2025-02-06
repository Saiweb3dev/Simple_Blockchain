package handlers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"Simple_Blockchain/models"
)

func NewTx(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create the transaction : %v", err)
		w.Write([]byte("Could not Create a new Transaction"))
		return
	}

	hash := md5.New()
	io.WriteString(hash, tx.Sender+tx.Receiver+string(tx.Amount))
	tx.ID = fmt.Sprintf("%x", hash.Sum(nil))

	resp, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal the transaction : %v", err)
		w.Write([]byte("Could not Create a new Transaction"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func WriteBlock(w http.ResponseWriter, r *http.Request){

	var checkoutTx models.TransactionCheckout
	if err := json.NewDecoder(r.Body).Decode(&checkoutTx); err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		log.Printf("Could not create the transaction checkout : %v", err)
		w.Write([]byte("Could not Create a new Transaction Checkout"))
	}
}
