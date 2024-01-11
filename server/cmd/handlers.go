package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"renting/internal/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // or use "*" to allow any origin
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
}

func (app *application) BuildingsReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.URL.Path != "/buildings" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	buildings, err := app.DB.GetBuildings()

	if err != nil {
		log.Println(err)
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buildings)
}

func (app *application) LoginReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.URL.Path != "/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var u user
	err := decoder.Decode(&u)

	lR := loginResoponce{"", models.User{}}

	if err != nil || u.Username == "" || u.Password == "" {
		lR.Status = "400"
		log.Println(err)
	} else {
		lR.User, err = app.DB.LoginUser(u.Username, u.Password)
		if err != nil {
			lR.Status = "400"
			log.Println(err)
		}

		lR.Status = "success"

		log.Println(lR)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lR)

}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResoponce struct {
	Status string      `json:"status"`
	User   models.User `json:"user"`
}

func (app *application) BuildingRegister(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add/building" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// read json payload

	var requestPayload struct {
		Title       string `json:"title"`
		RoomsNum    int    `json:"roomsNum"`
		ImageUrl    string `json:"imageUrl"`
		GuestsNum   int    `json:"guestsNum"`
		Description string `json:"description"`
		Country     string `json:"coutry"`
		City        string `json:"city"`
		Category    string `json:"category"`
		Address     string `json:"address"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		http.Error(w, "500 not found.", http.StatusBadRequest)
		return
	}

	
	fmt.Println(requestPayload)
}
