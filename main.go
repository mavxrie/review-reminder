package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func main() {
	pCustomGreeting := flag.String("greeting", "", "A custom greeting, only for jokes.")
	pJiraBrowseUrl := flag.String("jira-browse-url", "https://jira.com/browse/", "Your Jira browse url")
	pGithubToken := flag.String("github-token", "", "The github token to use (or use GITHUB_TOKEN)")
	pGithubOrg := flag.String("github-org", "avxSre", "The github organization to fetch PR from")
	pSlackTarget := flag.String("slack-target", "", "The slack target (channel, user) to send notification to.")
	pSlackToken := flag.String("slack-token", "", "The slack token to use (or use SLACK_TOKEN)")
	pHowMany := flag.Int("how-many", 3, "Number of PR to list.")
	flag.Parse()

	if *pGithubToken == "" {
		*pGithubToken = os.Getenv("GITHUB_TOKEN")
	}

	if *pSlackToken == "" {
		*pSlackToken = os.Getenv("SLACK_TOKEN")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *pGithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	results, _, err := client.Search.Issues(
		context.Background(),
		fmt.Sprintf("is:open is:pr archived:false org:%s draft:false review:required", *pGithubOrg),
		nil,
	)
	if err != nil {
		log.Fatalln(err.Error())
	}

	requests := make([]PullRequest, 0)

	for _, result := range results.Issues {
		parts := strings.Split(result.GetURL(), "/")

		reviewState, err := GetReviewState(client, parts[4], parts[5], result.GetNumber())
		if err != nil {
			log.Fatalln(err.Error())
		}

		requests = append(requests, PullRequest{
			ReviewState: reviewState,
			Org:         parts[4],
			Repository:  parts[5],
			Title:       result.GetTitle(),
			Number:      result.GetNumber(),
			Author:      result.GetUser().GetLogin(),
			CreatedAt:   result.GetCreatedAt(),
			UpdatedAt:   result.GetUpdatedAt(),
			Url:         result.PullRequestLinks.GetHTMLURL(),
		})
	}

	sort.SliceStable(requests, func(i, j int) bool {
		return requests[i].UpdatedAt.Before(requests[j].UpdatedAt)
	})

	out := ""
	greetings := GetGreetings(*pCustomGreeting)

	switch len(requests) {
	case 0:
		// no review
		return
	case 1:
		out = fmt.Sprintf("%s There is currenty one PR currently awaiting review:\n", greetings)
	default:
		out = fmt.Sprintf("%s There are currently %d PR currently awaiting reviews:\n", greetings, len(requests))
	}

	// Sort PR from oldest to newest
	for idx, result := range requests {
		if result.ReviewState == "approved" {
			continue
		}

		if *pSlackTarget != "" {
			out += result.SlackString(*pJiraBrowseUrl)
		} else {
			out += result.String()
		}

		if idx == *pHowMany-1 {
			break
		}
	}

	if len(requests) > *pHowMany {
		out += fmt.Sprintf("... and other can be viewed https://github.com/pulls/review-requested\n")
	}

	if *pSlackTarget != "" {
		SlackWrite(
			*pSlackToken,
			*pSlackTarget,
			out,
		)
	} else {
		fmt.Printf(out)
	}
}
