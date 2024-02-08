package main

import (
	"encoding/json"
	"net/http"

	"github.com/Scalingo/go-utils/logger"
)

var log = logger.Default()

func pongHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(map[string]string{"status": "pong"})
	if err != nil {
		log.WithError(err).Error("Fail to encode JSON")
	}
	return nil
}

type Owner struct {
	Login string `json:"login"`
}
type Repository struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Owner         Owner  `json:"owner"`
	Languages_URL string `json:"languages_url"`
	Languages     []string
}

func listGithubPublicRepositories(token string) ([]Repository, error) {
	url := "https://api.github.com/repositories"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	// Decode the JSON response into a slice of Repository structs
	//TODO: decode into interface to make the func reusable.
	var repositories []Repository
	err = json.NewDecoder(response.Body).Decode(&repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}
func getReposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())

	repositories, err := listGithubPublicRepositories(cfg.Token)
	if err != nil {
		log.WithError(err).Error("Fail to list repo JSON")
	}

	for _, repo := range repositories {
		log.Info("ID: %d, Name: %s, Full Name: %s\n", repo.ID, repo.Name)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(repositories)
	if err != nil {
		log.WithError(err).Error("Fail to encode JSON")
	}
	return nil
}

func getStatsHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//TODO: organize response
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Stats..."})
	if err != nil {
		log.WithError(err).Error("Fail to encode JSON")
	}
	return nil
}
