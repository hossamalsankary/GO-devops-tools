package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	Telnet "telnetapp/pkg/backend"
	Frontend "telnetapp/pkg/html"
)

func main() {

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/telnet", handleTelnet)

	fmt.Printf("Server starting on port 8080... \n")
	http.ListenAndServe(":8080", nil)

	// }
	//
	//	func telnetHandler(w http.ResponseWriter, r *http.Request) {
	//		Telnet.Telnet("127.0.0.1:80")
	//	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("telnetForm").Parse(Frontend.UI)
	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func handleTelnet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ip := r.FormValue("ip")
		port := r.FormValue("port")
		address := fmt.Sprintf("%s:%s", ip, port)
		// fmt.Fprintf(w, "Name: %s<br>IP: %s<br>Port: %s<br>", address, ip, port)
		fmt.Printf("Name: %s<br>IP: %s<br>Port: %s<br>", address, ip, port)
		telnetResponds, ok := Telnet.Telnet(address)

		// If connection is refused or any error occurs, display an alert to the user
		if !ok {

			fmt.Fprintf(w, `
					<html>
					<body>
						<script>alert("Failed to connect to %s \n Error: %s");</script>
						<a href="/">Go back</a>
					</body>
					</html>`, address, telnetResponds)
		} else {
			fmt.Fprintf(w, `
					<html>
					<body>
						<script>alert("successfully connected to %s \n message: %s");</script>
						<a href="/">Go back</a>
					</body>
					</html>`, address, telnetResponds)
		}
		return
	}

}
