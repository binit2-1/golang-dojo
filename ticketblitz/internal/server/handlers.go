package server

import (
	"encoding/json"
	"net/http"
	"ticketblitz/internal/domain"
)


type EventHandler struct{
	Repo domain.EventRepository
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request){
	var event domain.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err!=nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateEvent(&event)
	if err!=nil{
		http.Error(w, "Failed to create event in database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request){
	
	id := r.PathValue("id")

	event, err := h.Repo.GetEventByID(id)
	if err != nil{
		http.Error(w, "No Events Found", http.StatusNotFound)
		return
	}

	//ID doesn't exist in db
	if event == nil {
		http.Error(w, "Event Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}