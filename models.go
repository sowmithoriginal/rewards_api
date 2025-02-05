package main


// Receipt struct
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item struct
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ReceiptResponse struct
type ReceiptResponse struct {
	ID string `json:"id"`
}

// PointsResponse struct
type PointsResponse struct {
	Points int `json:"points"`
}
