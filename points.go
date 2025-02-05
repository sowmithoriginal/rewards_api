package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Compute points for a given receipt
func Calculate_Points(receipt Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

	// Converting total to float
	total, _ := strconv.ParseFloat(receipt.Total, 64)

	// 50 points if the total is a round dollar amount with no cents
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// If item description length is a multiple of 3, add price * 0.2 (rounded up)
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the purchase day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// 10 points if the purchase time is between 2:00 PM and 4:00 PM
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}