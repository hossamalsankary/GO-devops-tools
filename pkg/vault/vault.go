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

	host := strings.Split(r.Host, ":")[0]
	fmt.Printf("Host: %v\n", host)
	ip := os.Getenv("VAULT_IPS")
	fmt.Printf("IP: %v\n", ip)
	ipList := strings.Split(ip, ",")
	remoteAddr := r.RemoteAddr
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Printf(host, err)
	}

	// var ok bool

	// for i := range ipList {
	// 	if host == strings.TrimSpace(ipList[i]) {
	// 		ok = true

	// 	} else {

	// 		ok = false
	// 	}
	// }
	fmt.Printf("IP List: %v\n", ipList)
	fmt.Println(host)

	// if !ok {
	// 	http.Error(w, "Forbidden", http.StatusForbidden)
	// 	return
	// }
	// Cache the JSON content
	cachedJSON := Json.VaultJson

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cachedJSON)

}
