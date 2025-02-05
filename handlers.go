package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

// Process Receipt and generate an ID
func Process_Receipt_Handler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Generate a unique ID
	id := uuid.New().String()

	// Compute points and store in memory
	points := Calculate_Points(receipt)

	// Store receipt and points
	receiptStore[id] = receipt
	pointsStore[id] = points

	// Save to file
	save_Receipts_ToFile()

	log.Println("Stored receipt with ID:", id, "Points:", points)

	response := ReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Retrieve Points for a given receipt ID
func Get_Points_Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Received GET request for ID:", id)

	points, exists := pointsStore[id]
	if !exists {
		log.Println("ID not found in pointsStore:", id)
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}

	response := PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


