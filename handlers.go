package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

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

func getReposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())

	repositories, err := listGithubPublicRepositories(token)
	if err != nil {
		log.WithError(err).Error("Fail to list repo JSON")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	queryParams := r.URL.Query()
	paramValue := queryParams.Get("language")

	if paramValue != "" {
		var updatedRepos []*Repository
		for _, repo := range repositories {
			fmt.Println("repo.languages", repo.Languages)
			fmt.Println("paramValue", paramValue)
			for language := range repo.Languages {
				fmt.Println(language)

				if language == paramValue {
					updatedRepos = append(updatedRepos, repo)
				}
			}

		}
		err = json.NewEncoder(w).Encode(updatedRepos)
		if err != nil {
			log.WithError(err).Error("Fail to encode JSON")
		}
		return nil
	} else {
		err = json.NewEncoder(w).Encode(repositories)
		if err != nil {
			log.WithError(err).Error("Fail to encode JSON")
		}
	}

	return nil
}

func listGithubPublicRepositories(token string) ([]*Repository, error) {
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
	var repositories []*Repository
	err = json.NewDecoder(response.Body).Decode(&repositories)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for i := range repositories {
		wg.Add(1)
		go getLanguages(token, repositories[i].Languages_URL, &wg, repositories[i])
	}
	wg.Wait()

	fmt.Println(repositories)
	return repositories, nil
}

func getLanguages(token string, url string, wg *sync.WaitGroup, repo *Repository) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return
	}

	var languages map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&languages)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}
	repo.Languages = languages

}

func getAllReposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())

	repositories, err := listGithubPublicRepositories(token)
	if err != nil {
		log.WithError(err).Error("Fail to list repo JSON")
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
