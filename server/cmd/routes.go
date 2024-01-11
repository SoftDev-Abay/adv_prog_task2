package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/add/building", app.BuildingRegister)
	mux.HandleFunc("/login", app.LoginReciever)
	mux.HandleFunc("/buildings", app.BuildingsReciever)

	return mux
}
