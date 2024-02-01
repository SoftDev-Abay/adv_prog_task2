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
		Category:       requestPayload.Category,
	}
	fmt.Println(building)
	fmt.Println(app.DB.InsertBuilding(building))

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		lR.User, err = app.DB.LoginUser(u.Email, u.Password)
		if err == nil {
			lR.Status = "success"
		} else if err.Error() == "user not found" {
			lR.Status = "user not found"
			json.NewEncoder(w).Encode(lR)
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		} else if err.Error() == "invalid password" {
			lR.Status = "invalid password"
			json.NewEncoder(w).Encode(lR)
			http.Error(w, "invalid password", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
	if err != nil || u.Email == "" || u.Password == "" {
		aR.Status = "400"
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} else {
		exist, err := app.DB.UserExistsWithEmail(u.Email)
		if err != nil {
			http.Error(w, "db problem", http.StatusInternalServerError)
			return
		}

		if exist {
			fmt.Println("User is exist")
			aR.Status = "User is exist"
			json.NewEncoder(w).Encode(aR)
			// code 409 is conflict
			http.Error(w, "User is exist", http.StatusConflict)

			return
		}

		err = app.DB.Register(u)
		if err != nil {
			http.Error(w, "db problem", http.StatusInternalServerError)
			return
		}

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

// func (app *application) DeleteBuilding(w http.ResponseWriter, r *http.Request) {
// 	enableCors(&w)

// 	// Check for DELETE request method
// 	if r.Method != "DELETE" {
// 		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Extract building ID from URL
// 	vars := mux.Vars(r)
// 	buildingID := vars["id"]
// 	if buildingID == "" {
// 		http.Error(w, "Building ID is required", http.StatusBadRequest)
// 		return
// 	}

// 	// Call the database method to delete the building
// 	err := app.DB.
// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "Error deleting building", http.StatusInternalServerError)
// 		return
// 	}

// 	// Respond with a success message
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Building deleted successfully")
// }

// // Add the route in your main function or wherever you set up your routes
// // Assuming you are using gorilla/mux for routing
