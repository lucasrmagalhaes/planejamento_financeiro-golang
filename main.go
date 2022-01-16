package main

import (
	"net/http"

	"github.com/lucasrmagalhaes/planejamento_financeiro-golang/adapter/http/transaction"
)

func main() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.ListenAndServe(":8080", nil)
}
