package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type TransactionSummary struct {
	TotalTransactions  int                `json:"total_transactions"`
	TotalAmount        float64            `json:"total_amount"`
	ByType             map[string]float64 `json:"by_type"`
	ByStatus           map[string]int     `json:"by_status"`
	TransactionsPerDay map[string]int     `json:"transactions_per_day"`
}

func fetchAndLogSummary() {
	resp, err := http.Get(os.Getenv("BANK_API_URL"))
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
