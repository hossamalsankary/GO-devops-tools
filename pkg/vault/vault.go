package vault

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	Json "telnetapp/pkg/json"
)

func GetVaultToken(w http.ResponseWriter, r *http.Request) {

	// Restrictive access
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	ip := os.Getenv("VAULT_IPS")
	ipList := strings.Split(ip, ",")

	remoteAddr := r.RemoteAddr
	clientIp, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Printf(clientIp, err)
	}

	clientIp = strings.TrimSpace(clientIp)

	var ok bool = false

	for _, ip := range ipList {
		if clientIp == strings.TrimSpace(ip) {
			ok = true
			break
		}
	}

	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	// Cache the JSON content
	cachedJSON := Json.VaultJson

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cachedJSON)

}
