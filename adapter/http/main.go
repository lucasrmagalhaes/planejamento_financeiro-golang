package http

import (
	"net/http"

	"github.com/lucasrmagalhaes/planejamento_financeiro-golang/adapter/http/actuator"
	"github.com/lucasrmagalhaes/planejamento_financeiro-golang/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init Function
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
