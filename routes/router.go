package routes


import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/",getBlockchain).Methods("GET")
	r.HandleFunc("/",writeBlock).Methods("POST")
	r.HandleFunc("/newTx",newTx).Methods("POST")
	return r
}