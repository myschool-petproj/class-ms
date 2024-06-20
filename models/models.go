package models

type Profile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	Id        string `json:"id"`
	Grade     int    `json:"grade"`
	ProfileID string `json:"profile_id"`
}
