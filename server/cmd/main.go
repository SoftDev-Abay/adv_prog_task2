package main

import (
	"fmt"
	"log"
	"net/http"
	"renting/internal/repository"
	"renting/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	DB repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// connect to DB

	conn := app.GetDBInstance()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	// defer conn.Close()

	// start a web server
	fmt.Println("Stargin app on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
