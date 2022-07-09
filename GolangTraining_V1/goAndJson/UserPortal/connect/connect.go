package connect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	connectDB := "user=postgres dbname=postgres password=Pass@123 host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connectDB) //This postgres is the driver
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
	return db
}
