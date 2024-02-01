package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"renting/internal/service"
	"renting/models"
)

func (h *Handlers) BuildingsReciever(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	buildings, err := h.c.Store.GetBuildings()

	if err != nil {
		log.Println(err)
		http.Error(w, "500", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buildings)
}

func (h *Handlers) BuildingRegister(w http.ResponseWriter, r *http.Request) {
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
	err := service.ReadJson(w, r, &requestPayload)
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
	fmt.Println(h.c.Store.InsertBuilding(building))

}
