package routes

import (
	"Simple_Blockchain/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// r.HandleFunc("/",getBlockchain).Methods("GET")
	// r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/newTx", handlers.NewTx).Methods("POST")
	return r
}
