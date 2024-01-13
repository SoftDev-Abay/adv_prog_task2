package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"renting/models"
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

	building := models.Building{
		Description:    requestPayload.Description,
		Address:        requestPayload.Address,
		Country:        requestPayload.Country,
		GuestsNum:      4,
		RoomsNum:       2,
		BathroomsNum:   2,
		PriceDay:       100,
		AvalableFrom:   "2024-01-01",
		AvalableUntill: "2024-01-10",
		UserId:         1,
		ImgUrl:         requestPayload.ImageUrl,
		City:           requestPayload.City,
		Category:       1,
	}
	id, err := app.DB.InsertBuilding(building)
	fmt.Println(id, err)
	fmt.Println(requestPayload)
}

func (app *application) LoginReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.URL.Path != "/auth/login" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var u models.User
	err := decoder.Decode(&u)

	lR := AuthResoponce{"", models.User{}}

	if err != nil || u.Email == "" || u.Password == "" {
		lR.Status = "400"
		log.Println(err)
	} else {
		lR.User, err = app.DB.LoginUser(u.Email, u.Password)
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

func (app *application) RegisterReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.URL.Path != "/auth/signup" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var u models.User
	err := decoder.Decode(&u)

	aR := AuthResoponce{"", u}
	fmt.Println(u)
	if err != nil || u.Email == "" || u.Password == "" {
		aR.Status = "400"
		log.Println(err)
	} else {
		// err = app.DB.Register(u)
		// if err != nil {
		// 	aR.Status = "400"
		// 	log.Println(err)
		// }

		aR.Status = "success"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aR)
}

type AuthResoponce struct {
	Status string      `json:"status"`
	User   models.User `json:"user"`
}
