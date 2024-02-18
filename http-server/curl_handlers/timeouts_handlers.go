package curl_handlers

import (
	"Distributed-calculator-of-arithmetic-expressions/internal/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func SetSettingHandler(w http.ResponseWriter, r *http.Request) {
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

	timeout := values.Get("setting")
	if timeout == "" {
		http.Error(w, "Missing 'timeout' parameter in request body", http.StatusBadRequest)
		return
	}

	if err := storage.SetNewSetting(timeout); err != nil {
		http.Error(w, "Could not set new expression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Timeout added successfully"))
}

func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	timeouts, err := storage.GetSettings()
	if err != nil {
		http.Error(w, "Failed to retrieve expressions"+err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(timeouts)
	if err != nil {
		http.Error(w, "Failed to marshal expressions", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
