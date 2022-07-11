package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/xeonx/timeago"
)

type PullRequest struct {
	ReviewState string
	Org         string
	Repository  string
	Title       string
	Number      int
	Author      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Url         string
}

func (pr PullRequest) String() string {
	cfg := timeago.NoMax(timeago.English)

	return fmt.Sprintf(
		"%s/%s#%d by %s: %s updated %s - %s\n",
		pr.Org,
		pr.Repository,
		pr.Number,
		pr.Author,
		pr.Title,
		cfg.Format(pr.UpdatedAt),
		pr.Url,
	)
}

func (pr PullRequest) SlackString(jiraBrowseUrl string) string {
	cfg := timeago.NoMax(timeago.English)

	title := pr.Title

	re := regexp.MustCompile(`AVX[A-Z]*-[0-9]*`)
	jiraID := re.Find([]byte(title))

	title = strings.Replace(title, string(jiraID), fmt.Sprintf("<%s%s|%s>", jiraBrowseUrl, string(jiraID), string(jiraID)), -1)

	return fmt.Sprintf(
		"<%s|%s/%s#%d> by %s: %s updated %s\n",
		pr.Url,
		pr.Org,
		pr.Repository,
		pr.Number,
		pr.Author,
		title,
		cfg.Format(pr.UpdatedAt),
	)
}
