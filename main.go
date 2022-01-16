package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lucasrmagalhaes/planejamento_financeiro-golang/model/transaction"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createTransaction)

	http.ListenAndServe(":8080", nil)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2022-01-16T10:20:00")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    3000.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
