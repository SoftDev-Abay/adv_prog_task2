package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "db_renting"
)

var (
	DB     *sql.DB
	once   sync.Once
	dbOnce sync.Once
)

func getDB() *sql.DB {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			HOST, PORT, USER, PASSWORD, DBNAME)

		dbOnce.Do(func() {
			var err error
			DB, err = sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}

			err = DB.Ping()
			if err != nil {
				panic(err)
			}
		})
	})

	return DB
}

func (app *application) GetDBInstance() *sql.DB {
	fmt.Println("connected to DB")
	return getDB()
}
