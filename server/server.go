package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"renting/classes"
	db "renting/database"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// func registerHandler(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path != "/form" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "POST" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")

// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResoponce struct {
	Status string       `json:"status"`
	User   classes.User `json:"user"`
}

func loginReciever(w http.ResponseWriter, r *http.Request) {

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

	lR := loginResoponce{"", classes.User{}}

	if err != nil || u.Username == "" || u.Password == "" {
		lR.Status = "400"
		log.Println(err)
	} else {
		lR.User, err = db.LoginUser(u.Username, u.Password)
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

// func registerReciever(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path != "/register" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "POST" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)

// 	var u user
// 	err := decoder.Decode(&u)

// 	if err != nil || u.Username == "" || u.Password == "" {
// 		log.Println(err)
// 	} else {
// 		err = db.Register(u.Username, u.Password, u.Email, u.PhoneNum)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(u)
// }

func BuildingsReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.URL.Path != "/buildings" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	buildings, err := db.GetBuildings()

	if err != nil {
		log.Println(err)
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buildings)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	// http.HandleFunc("/form", registerHandler)
	http.HandleFunc("/login", loginReciever)
	http.HandleFunc("/buildings", BuildingsReciever)

	fmt.Printf("Starting live reloaded server at port 8080\n")

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}

}
