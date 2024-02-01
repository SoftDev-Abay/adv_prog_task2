package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()

	router.Use(AllowOptionsMiddleware)

	router.HandleFunc("/add/building", s.h.BuildingRegister)
	router.HandleFunc("/auth/login", s.h.LoginReciever)
	router.HandleFunc("/auth/signup", s.h.RegisterReciever)
	router.HandleFunc("/buildings", s.h.BuildingsReciever).Methods("GET")
	// mux.HandleFunc("/delete/building/{id}", app.DeleteBuilding).Methods("DELETE")

	return router
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
