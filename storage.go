package main



import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Global storage variables
var (
	receiptStore = make(map[string]Receipt) // Stores receipts with their IDs
	pointsStore  = make(map[string]int)     // Stores computed points per receipt ID
)

const fileName = "receipts.json"


// Save receipts to a file
//logging to catch any file related errors
func save_Receipts_ToFile() {
	data, err := json.Marshal(receiptStore)
	if err != nil {
		log.Println("Error saving receipts:", err)
		return
	}
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

// Load receipts from a file when the server starts
//logging to catch any file realted errors
func load_Receipts_FromFile() {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return // No file exists yet
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Error reading receipts file:", err)
		return
	}

	err = json.Unmarshal(data, &receiptStore)
	if err != nil {
		log.Println("Error decoding receipts file:", err)
	}
}
