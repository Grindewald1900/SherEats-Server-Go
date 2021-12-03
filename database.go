package main

/*
Some database operations, CRUD
*/
import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/lib/pq"
)

// var database *sql.DB

// Postgre basic information(local)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "ShereatsDB"
)

// Insert
func InsertDatabase() {
	// println("ConnectDatabase")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// connect database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	println(reflect.TypeOf(db))

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// println("Successfully connected!")
	// println("InsertDatabase")

	sqlStatement := `
	INSERT INTO restaurant(id, name, address, genre, averageprice, tel, lat, long, isfavourite)
	VALUES(5, 'AAA', 'BBB', 'CCC', 12, '8190000000', 78.0011, 23.1111, true)
	RETURNING id`
	id := 0
	println(sqlStatement)
	err2 := db.QueryRow(sqlStatement).Scan(&id)
	if err2 != nil {
		panic(err2)
	}
	println("Insert connected id is:", id)
}
