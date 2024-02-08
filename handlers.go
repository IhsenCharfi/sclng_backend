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
	ch := make(chan interface{}, len(repositories))

	for i := range repositories {
		wg.Add(1)
		go getLanguages(token, repositories[i].Languages_URL, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for _, repo := range repositories {
		languages := <-ch
		if languages != nil {
			repo.Languages = languages
		}
		fmt.Println(*repo)
	}
	fmt.Println(repositories)
	return repositories, nil
}

func getLanguages(token string, url string, wg *sync.WaitGroup, ch chan<- interface{}) {
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
		ch <- nil
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return
	}

	var languages interface{}
	err = json.NewDecoder(response.Body).Decode(&languages)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		ch <- nil
		return
	}

	ch <- languages
}

func getReposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())

	repositories, err := listGithubPublicRepositories(token)
	if err != nil {
		log.WithError(err).Error("Fail to list repo JSON")
	}
	for _, repo := range repositories {
		fmt.Println(repo)
		fmt.Println(*repo)

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
