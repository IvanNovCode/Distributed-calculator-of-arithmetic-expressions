package main

import (
	"Distributed-calculator-of-arithmetic-expressions/http-server/agent_handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/addexpr", agent_handlers.AddExprHandler)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		return
	}
}
