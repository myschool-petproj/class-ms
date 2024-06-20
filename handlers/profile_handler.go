package handlers

import (
	"encoding/json"
	"net/http"
)

func GetProfiles(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("All profiles")
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Profile")
}
