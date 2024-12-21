package vault

import (
	"encoding/json"
	"net/http"
	Json "telnetapp/pkg/json"
)

func GetVaultToken(w http.ResponseWriter, r *http.Request) {
	// Cache the JSON content
	cachedJSON := Json.VaultJson

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cachedJSON)

}
