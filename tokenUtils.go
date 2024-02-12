package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func getClient(token string) *github.Client {
	//create github client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return client
}

func extractToken(r *http.Request) string {
	// Retrieve the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is not present
	if authHeader == "" {
		return "No token provided"
	}

	// Split the header value to get the token part
	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		return "Invalid Authorization header format"
	}

	return authParts[1]
}

func isGitHubTokenValid(token string) bool {

	client := getClient(token)
	// Make a request to the authenticated user endpoint to check token validity
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		log.WithError(err).Error("Error checking token validity")
		return false
	}

	// If the request was successful, the token is valid
	log.Info("Authenticated user: %s\n", user.GetLogin())
	return true
}
