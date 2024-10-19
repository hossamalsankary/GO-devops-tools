package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telnetapp/pkg/curl"
	Database "telnetapp/pkg/database"
	Telnet "telnetapp/pkg/telnet"
)

func main() {

	// (username string, password string, dbname string, host string)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/telnet", handleTelnet)
	http.HandleFunc("/curl", handleCurl)
	http.HandleFunc("/db", handleDbConnection)

	fmt.Printf("Server starting on port 8080... \n")
	http.ListenAndServe(":8080", nil)

}

func handleCurl(w http.ResponseWriter, r *http.Request) {
	curl.HandleCurl(w, r)
}

func handleTelnet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Define a struct to hold the incoming JSON data
		var requestData struct {
			IP   string `json:"ip"`
			Port string `json:"port"`
		}

		// Decode the JSON data from the request body
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Use the decoded data
		address := fmt.Sprintf("%s:%s", requestData.IP, requestData.Port)

		telnetResponds, ok := Telnet.Telnet(address)
		response := map[string]interface{}{
			"success": ok,
			"message": telnetResponds,
		}

		// Set the response header and write the JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// If the method is not POST, return an error
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func handleDbConnection(w http.ResponseWriter, r *http.Request) {

	Database.CheckDbconnection(w, r)
}
