package curl

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type CurlRequest struct {
	Command string `json:"command"`
}

type CurlResponse struct {
	Output  string `json:"content"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

func HandleCurl(w http.ResponseWriter, r *http.Request) {
    var req CurlRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    cmd := exec.Command("sh", "-c", req.Command)
    output, err := cmd.CombinedOutput()

    response := CurlResponse{
        Output:  string(output),
        Success: err == nil,
    }
    if err != nil {
        response.Error = err.Error()
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}