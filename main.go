package main

import (
	"net/http"
	"log"

	"Simple_Blockchain/routes"
)


func main(){

	r := routes.Router()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}