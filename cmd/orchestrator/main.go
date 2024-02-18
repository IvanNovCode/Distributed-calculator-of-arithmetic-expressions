package main

import (
	"Distributed-calculator-of-arithmetic-expressions/http-server/curl_handlers"
	"Distributed-calculator-of-arithmetic-expressions/http-server/orchestrator_handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getexpressions", curl_handlers.GetExpressionsHandler)
	mux.HandleFunc("/addexpression", curl_handlers.AddExpressionHandler)
	mux.HandleFunc("/clearexpressions", curl_handlers.ClearExpressionsHandler)
	mux.HandleFunc("/setsetting", curl_handlers.SetSettingHandler)
	mux.HandleFunc("/getsettings", curl_handlers.GetSettingsHandler)
	mux.HandleFunc("/getanswer", orchestrator_handlers.GetAnswerHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
