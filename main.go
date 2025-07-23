package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TransactionSummary struct {
	TotalTransactions  int                `json:"total_transactions"`
	TotalAmount        float64            `json:"total_amount"`
	ByType             map[string]float64 `json:"by_type"`
	ByStatus           map[string]int     `json:"by_status"`
	TransactionsPerDay map[string]int     `json:"transactions_per_day"`
}

func fetchAndLogSummary() {
	serviceURL := os.Getenv("CHOREO_SUMMARY_CONN_SERVICEURL")
	choreoApiKey := os.Getenv("CHOREO_SUMMARY_CONN_APIKEY")

	// Create a new HTTP request
	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Add API key to the header (adjust the header name if required by your API)
	req.Header.Set("Choreo-API-Key", fmt.Sprintf("Bearer %s", choreoApiKey))
	// OR use a custom header if needed:
	// req.Header.Set("X-API-Key", choreoApiKey)

	// Send the request using http.Client
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch transaction summary: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API returned non-OK status: %v", resp.Status)
	}

	var summary TransactionSummary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	log.Println("🔔 Transaction Summary Report")
	log.Printf("📊 Total Transactions: %d\n", summary.TotalTransactions)
	log.Printf("💰 Total Amount: %.2f\n", summary.TotalAmount)

	log.Println("📂 Breakdown by Transaction Type:")
	for typ, amt := range summary.ByType {
		log.Printf("  - %s: %.2f", typ, amt)
	}

	log.Println("📁 Breakdown by Status:")
	for status, count := range summary.ByStatus {
		log.Printf("  - %s: %d", status, count)
	}

	log.Println("🗓️ Transactions Per Day:")
	for day, count := range summary.TransactionsPerDay {
		log.Printf("  - %s: %d", day, count)
	}
}

func main() {
	log.Println("Starting Transaction Summary Logger...")

	fetchAndLogSummary()

}
