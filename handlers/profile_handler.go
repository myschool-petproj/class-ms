package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/myschool-petproj/class-ms/service"
	"net/http"
)

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	profiles := service.GetAllProfiles()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profiles)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := uuid.Validate(id)
	if err != nil {
		http.Error(w, "id should be uuid", http.StatusBadRequest)
		return
	}
	profile, err := service.GetProfileById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(profile)
}
