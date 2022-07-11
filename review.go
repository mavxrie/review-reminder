package main

import (
	"context"

	"github.com/google/go-github/v45/github"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// GetReviewState checks the PR state: open, changes requested, etc...
func GetReviewState(client *github.Client, org, repository string, issue int) (string, error) {
	reviews, _, err := client.PullRequests.ListReviews(
		context.Background(),
		org,
		repository,
		issue,
		nil,
	)

	if err != nil {
		return "", err
	}

	states := make([]string, 0)

	for _, review := range reviews {
		if review.GetState() == "COMMENTED" {
			continue
		}

		if review.GetState() == "DISMISSED" {
			continue
		}

		states = append(states, *review.State)
	}

	if contains(states, "APPROVED") {
		return "approved", nil
	} else if contains(states, "CHANGES_REQUESTED") {
		return "changes_requested", nil
	}

	return "open", nil
}
