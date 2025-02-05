package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func main() {
	// Load stored receipts on startup
	load_Receipts_FromFile()

	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", Process_Receipt_Handler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", Get_Points_Handler).Methods("GET")

	log.Println("Server running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", r))
}
