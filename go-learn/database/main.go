package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "eugenewong",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my_album",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	var albums []Album

	result, err := db.Query("SELECT * FROM albumDB")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var alb Album
		if err := result.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			fmt.Println(err.Error())
		}
		albums = append(albums, alb)
	}
	if err := result.Err(); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(albums)
}
