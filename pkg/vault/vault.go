package vault

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetVaultToken(w http.ResponseWriter, r *http.Request) {
	// open the file
	file, err := os.Open("app/dummy.json")

	if err != nil {
		panic(err)
	}

	// read the file
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// Cache the JSON content
	cachedJSON := data
	fmt.Println(r.URL.Path, r.Host)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cachedJSON)

}
