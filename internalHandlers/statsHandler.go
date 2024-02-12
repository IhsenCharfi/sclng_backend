package internalHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/IhsenCharfi/sclng_backend/models"
	"github.com/IhsenCharfi/sclng_backend/utils"

	"github.com/Scalingo/go-utils/logger"
)

func GetStatsHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
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

	//get filter param
	queryParams := r.URL.Query()
	paramValue := queryParams.Get("language")
	url := "http://localhost:3000/repos"
	if paramValue != "" {
		url = url + "?language=" + paramValue
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil
	}

	var repositories []*models.Repository
	err = json.NewDecoder(response.Body).Decode(&repositories)
	if err != nil {
		return nil
	}

	stats := map[string]interface{}{
		"total_repos": len(repositories),
		//"largest_bytes": largestBytes(repositories),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return nil
	}
	return nil
}

func largest_bytes(resporitories []*models.Repository) {
	//max := 0
	//for _, repo := range resporitories {
	//TODO..
	//}
}
