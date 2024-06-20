package handlers

import (
	"encoding/json"
	"net/http"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("All Classes")
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Class")
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Create Class")
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Update Class")
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode("Delete class")
}
