package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "localhost\\SQLExpress"
	port     = 1433
	user     = "sa"
	password = "mack2967"
	database = "sca"
)

func readEmployees(db *sql.DB, w http.ResponseWriter) (int, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		return -1, err
	}
	defer conn.Close()

	tsql := fmt.Sprintf("SELECT Id, LastName, FirstName FROM dbo.Employee;")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var lastname, firstname string
		var id int
		err := rows.Scan(&id, &lastname, &firstname)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Fprintf(w, "ID: %d, LastName: %s, FirstName: %s\n", id, lastname, firstname)
		count++
	}
	return count, nil
}
