package main

import (
	"assignment_3/pkg/store/postgres"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("DB_PORT")
	user := "web"
	password := "admin"
	dbname := "library"
	db, err := postgres.OpenDB(port, user, password, dbname)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

}
