package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/myschool-petproj/class-ms/models"
	"sync"
)

var (
	profiles = map[string]*models.Profile{}
	mutex    = &sync.Mutex{}
)

func init() {
	mutex.Lock()
	defer mutex.Unlock()

	defaultProfiles := []*models.Profile{
		{Id: uuid.New().String(), Name: "John Doe"},
		{Id: uuid.New().String(), Name: "Jane Smith"},
	}

	for _, profile := range defaultProfiles {
		profiles[profile.Id] = profile
	}
}

func GetAllProfiles() []*models.Profile {
	mutex.Lock()
	defer mutex.Unlock()

	var allProfiles []*models.Profile
	for _, profile := range profiles {
		allProfiles = append(allProfiles, profile)
	}
	return allProfiles
}

func GetProfileById(id string) (*models.Profile, error) {
	mutex.Lock()
	defer mutex.Unlock()

	profile, exists := profiles[id]
	if !exists {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}
