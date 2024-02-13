package internalHandlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/IhsenCharfi/sclng_backend/models"
	"github.com/IhsenCharfi/sclng_backend/utils"

	"github.com/Scalingo/go-utils/logger"
)

var log = logger.Default()

func GetReposHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
	log := logger.Get(r.Context())

	//Check token validation
	token := utils.ExtractToken(r)
	valid := utils.IsGitHubTokenValid(token)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "token invalid"})
		if err != nil {
			log.WithError(err).Error("Failed to encode JSON")
		}
		return err
	}

	repositories, err := listGithubPublicRepositories(token)
	if err != nil {
		log.WithError(err).Error("Fail to list repo JSON")
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
		if err != nil {
			log.WithError(err).Error("Fail to list repo JSON")
		}
		return err
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	queryParams := r.URL.Query()
	paramValue := queryParams.Get("language")

	if paramValue != "" {
		var updatedRepos []*models.Repository
		for _, repo := range repositories {
			for language := range repo.Languages {
				if language == paramValue {
					updatedRepos = append(updatedRepos, repo)
				}
			}

		}
		err = json.NewEncoder(w).Encode(updatedRepos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
			if err != nil {
				log.WithError(err).Error("Fail to encode JSON")
			}
			return err
		}
		return nil
	} else {
		err = json.NewEncoder(w).Encode(repositories)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(map[string]string{"error": "internal error"})
			if err != nil {
				log.WithError(err).Error("Fail to encode JSON")
			}
			return err
		}
	}

	return nil
}

func listGithubPublicRepositories(token string) ([]*models.Repository, error) {
	url := "https://api.github.com/search/repositories?apiVersion=2022-11-28&q=is:public&sort=updated&order=desc&per_page=100"
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
		// TODO: fix it when status not OK and err nil
		return nil, err
	}

	// Decode the JSON response into a slice of SearchResponse structs
	var repoResponse models.SearchResponse
	err = json.NewDecoder(response.Body).Decode(&repoResponse)
	if err != nil {
		return nil, err
	}

	var repositories []*models.Repository
	repositories = repoResponse.Items

	var wg sync.WaitGroup

	for i := range repositories {
		wg.Add(1)
		go getLanguages(token, repositories[i].Languages_URL, &wg, repositories[i])
	}
	wg.Wait()

	return repositories, nil
}

func getLanguages(token string, url string, wg *sync.WaitGroup, repo *models.Repository) {
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

	var languages map[string]int
	err = json.NewDecoder(response.Body).Decode(&languages)
	if err != nil {
		log.WithError(err).Error("Fail to decode JSON")
		return
	}
	repo.Languages = languages

}
