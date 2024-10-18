package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"telnetapp/pkg/curl"
	Database "telnetapp/pkg/database"
	Telnet "telnetapp/pkg/telnet"
)

func main() {
	// (username string, password string, dbname string, host string)
	Database.InitSQLDB("myuser", "dumy", "mydb", "127.0.0.1")
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

		ip := r.FormValue("ip")
		port := r.FormValue("port")
		address := fmt.Sprintf("%s:%s", ip, port)
		// fmt.Fprintf(w, "Name: %s<br>IP: %s<br>Port: %s<br>", address, ip, port)
		//fmt.Printf("Name: %s IP: %s Port: %s", address, ip, port)
		telnetResponds, ok := Telnet.Telnet(address)
		response := map[string]interface{}{
			"success": ok,
			"message": telnetResponds,
		}
		responseJsonData, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		// If connection is refused or any error occurs, display an alert to the user
		if !ok {

			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJsonData)

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(responseJsonData)
		}
		return
	}

}

func handleDbConnection(w http.ResponseWriter, r *http.Request) {

	Database.InitSQLDB("myuser", "dumy", "mydb", "127.0.0.1")
}
