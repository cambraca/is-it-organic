package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
)

const postgresConn = "host=postgres port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", postgresConn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	token := flag.String("token", "", "twitter API token")
	query := flag.String("query", "", "twitter query")
	flag.Parse()

	search(*token, *query)
}
