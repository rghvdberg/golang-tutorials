package main

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties
	// We use the credentials from the environment variables
	// in the docker-compose file.
	// DBUSER = root
	// DBPASS = my-secret-pw
	cfg := mysql.Config{
		//User:  os.Getenv("DBUSER"),
		//Passwd: os.Getenv("DBPASS"),
		User:                 "root",
		Passwd:               "my-secret-pw",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
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
	fmt.Println("Connected to database")
}
