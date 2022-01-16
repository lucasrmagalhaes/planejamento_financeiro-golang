package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8080", nil)
}

type Transaction struct {
	Title     string
	Amount    float32
	Type      int
	CreatedAt time.Time
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2022-01-16T10:20:00")

	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    3000.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	json.NewEncoder(w).Encode(transactions)
}
