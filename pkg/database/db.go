package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSQLDB(username string, password string, dbname string, host string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	connectionErorr := db.Ping()

	if connectionErorr != nil {
		fmt.Printf("Database connection failed: %v", connectionErorr)
	} else {
		fmt.Println("Database connection successful")
	}
}
