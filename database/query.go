package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://jeoysqgj:Yqpkng49GujaIUn9LrGfzP1bRD3JHFAM@suleiman.db.elephantsql.com:5432/jeoysqgj")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one row statement", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)
	var id int
	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't scan roe into variables", err)
	}

	fmt.Println("one row", id, title, status)
}
