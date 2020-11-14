package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", "postgres://jeoysqgj:Yqpkng49GujaIUn9LrGfzP1bRD3JHFAM@suleiman.db.elephantsql.com:5432/jeoysqgj")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")

	fmt.Println("okay")
}
