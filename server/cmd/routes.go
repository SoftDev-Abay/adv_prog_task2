package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/add/building", app.BuildingRegister)
	mux.HandleFunc("/auth/login", app.LoginReciever)
	mux.HandleFunc("/auth/signup", app.RegisterReciever)
	mux.HandleFunc("/buildings", app.BuildingsReciever)

	return AllowOptionsMiddleware(mux)
}
func AllowOptionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			// Handle OPTIONS request and allow everything
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return
		}

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
