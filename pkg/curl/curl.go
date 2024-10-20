package curl

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request struct represents the incoming request payload
type Request struct {
	URL     string  `json:"url"`
	Method  string  `json:"method"` // GET, POST, etc.
	Data    string  `json:"data"`   // For POST, PUT requests
	Timeout float64 `json:"timeout,omitempty"`
	Proxy   string  `json:"proxy,omitempty"`
}

// Response struct represents the response payload sent back to the client
type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func HandleCurl(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendResponse(w, Response{
			Message: "Invalid request body",
			Success: false,
		})
		return
	}

	client := &http.Client{Timeout: 30 * time.Second}

	if req.Timeout > 0 {
		client.Timeout = time.Duration(req.Timeout * float64(time.Second))
	}

	if req.Proxy != "" {
		proxyURL, err := url.Parse(req.Proxy)
		if err != nil {
			sendResponse(w, Response{
				Message: "Invalid proxy URL",
				Success: false,
			})
			return
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	// Create the HTTP request
	httpReq, err := http.NewRequest(req.Method, req.URL, strings.NewReader(req.Data))
	if err != nil {
		sendResponse(w, Response{
			Message: "Failed to create HTTP request: " + err.Error(),
			Success: false,
		})
		return
	}

	// Perform the HTTP request
	resp, err := client.Do(httpReq)
	if err != nil {
		sendResponse(w, Response{
			Message: "HTTP request failed: " + err.Error(),
			Success: false,
		})
		return
	}
	defer resp.Body.Close()

	// At this point, the request was successful
	sendResponse(w, Response{
		Message: "HTTP request successful",
		Success: true,
	})
}

func sendResponse(w http.ResponseWriter, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
