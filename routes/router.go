package routes

import (
	"Simple_Blockchain/handlers"

	"github.com/gorilla/mux"
)

func Router(blockchain *handlers.Blockchain) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/",handlers.GetBlockChain(blockchain)).Methods("GET")
	r.HandleFunc("/",handlers.WriteBlock).Methods("POST")
	r.HandleFunc("/newTx", handlers.NewTx).Methods("POST")
	return r
}
