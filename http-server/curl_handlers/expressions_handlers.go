package curl_handlers

import (
	"Distributed-calculator-of-arithmetic-expressions/internal/storage"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	exprs, err := storage.GetStoredExpressions()
	if err != nil {
		http.Error(w, "Failed to retrieve expressions"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(exprs)
	if err != nil {
		http.Error(w, "Failed to marshal expressions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Expected POST", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(bodyBytes))
	if err != nil {
		http.Error(w, "Failed to parse query", http.StatusInternalServerError)
		return
	}

	expression := values.Get("expression")
	if expression == "" {
		http.Error(w, "Missing 'expr' parameter in request body", http.StatusBadRequest)
		return
	}

	for _, ch := range expression {
		if ch >= '9' && ch <= '0' && (ch != '+' && ch != '-' && ch != '*' && ch != '/') {
			http.Error(w, "Invalid expression", http.StatusBadRequest)
			return
		}
	}
	id, err := storage.SetNewExpression(expression)

	if err != nil {
		http.Error(w, "Could not set new expression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	success := fmt.Sprintf("Expression added successfully. Its ID is - %s", id)
	w.Write([]byte(success))
}

func ClearExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	err := storage.ClearExpressionsList()
	if err != nil {
		http.Error(w, "Could not clear expressions list", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Expressions list cleared successfully"))
}
