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

	defer db.Close()

	stmt, err := db.Prepare("UPDATE todos SET status=$2 where id=$1")

	if err != nil {
		log.Fatal("can't prepare statement update")
	}

	if _, err := stmt.Exec(1, "inactive"); err != nil {
		log.Fatal("error execite update ", err)
	}

	fmt.Println("update success")
}
