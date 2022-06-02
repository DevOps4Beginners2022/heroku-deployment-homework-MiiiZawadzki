package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello DaftAcademy!")
	})

	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		secret := os.Getenv("SECRET")
		if secret == "" {
			fmt.Fprintf(w, "$SECRET not set")
			return
		}
		fmt.Fprintf(w, secret)
	})

	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
		db_url := os.Getenv("DATABASE_URL")

		if db_url == "" {
			fmt.Fprintf(w, "$DATABASE_URL not set")
			return
		}
		// Connect to db
		db, err := sql.Open("postgres", db_url)

		if err != nil {
			fmt.Fprintf(w, "Error! Cannot connect to database")
			log.Fatal(err)
			return
		}

		fmt.Fprintf(w, "Database is ready to use!")
		// close database
		defer db.Close()

	})

	http.ListenAndServe(":"+port, nil)
}
