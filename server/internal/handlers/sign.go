package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"renting/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // or use "*" to allow any origin
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token, Authorization")
}

func (h *Handlers) LoginReciever(w http.ResponseWriter, r *http.Request) {
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
		lR.User, err = h.c.Store.LoginUser(u.Email, u.Password)
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

func (h *Handlers) RegisterReciever(w http.ResponseWriter, r *http.Request) {
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
		exist, err := h.c.Store.UserExistsWithEmail(u.Email)
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

		err = h.c.Store.Register(u)
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
