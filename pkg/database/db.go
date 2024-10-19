package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type dbConnection struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
}
type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func InitSQLDB(username string, password string, dbname string, host string, port int) (string, bool) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	connectionErorr := db.Ping()

	if connectionErorr != nil {
		return fmt.Sprintf("Database connection failed: %v", connectionErorr), false
	} else {
		return "Database connection successful", true
	}
}

func CheckDbconnection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req dbConnection
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbResponse, ok := InitSQLDB(req.Username, req.Password, req.Dbname, req.Host, req.Port)

	response := response{
		Success: ok,
		Message: dbResponse,
	}
	fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
