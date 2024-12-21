package vault

import (
	"encoding/json"
	"fmt"
	"net/http"
	Json "telnetapp/pkg/json"
)

func GetVaultToken(w http.ResponseWriter, r *http.Request) {
	// Cache the JSON content
	cachedJSON := Json.VaultJson

	w.Header().Set("Content-Type", "application/json")
	fmt.Printf(r.Host)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cachedJSON)

}
