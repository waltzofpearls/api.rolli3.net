package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type (
	githubRepo struct {
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
		Fullname    *string `json:"fullname,omitempty"`
		Branch      *string `json:"branch,omitempty"`
		Url         *string `json:"url,omitempty"`
		Language    *string `json:"language,omitempty"`
		Forks       *int    `json:"forks,omitempty"`
		Stars       *int    `json:"stars,omitempty"`
		Watches     *int    `json:"watches,omitempty"`
		Badge       *string `json:"badge,omitempty"`
	}

	githubContribs           map[string]int
	getGithubContribsDocFunc func(*config) (*goquery.Document, error)

	githubAPI struct {
		client    *github.Client
		docGetter getGithubContribsDocFunc
		config    *config
	}
)

func getGithubContribsDoc(c *config) (*goquery.Document, error) {
	url := fmt.Sprintf(
		"https://github.com/users/%s/contributions",
		c.GithubUsername,
	)
	return goquery.NewDocument(url)
}

func newGithubAPI(c *config) *githubAPI {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: c.GithubToken,
		},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &githubAPI{
		client:    github.NewClient(tc),
		docGetter: getGithubContribsDoc,
		config:    c,
	}
}

func (g *githubAPI) getRepos() (*[]githubRepo, error) {
	opt := &github.RepositoryListOptions{
		Type:      "all",
		Sort:      "pushed",
		Direction: "desc",
	}
	repos, _, err := g.client.Repositories.List("", opt)
	if err != nil {
		return nil, err
	}

	var gr []githubRepo
	for _, repo := range repos {
		badge := fmt.Sprintf(
			"https://api.travis-ci.org/%s.svg?branch=%s",
			*repo.FullName,
			*repo.DefaultBranch,
		)
		log.Printf("%#v\n\n", repo)
		gr = append(gr, githubRepo{
			Name:        repo.Name,
			Description: repo.Description,
			Fullname:    repo.FullName,
			Branch:      repo.DefaultBranch,
			Url:         repo.HTMLURL,
			Language:    repo.Language,
			Forks:       repo.ForksCount,
			Stars:       repo.StargazersCount,
			Watches:     repo.SubscribersCount,
			Badge:       &badge,
		})
	}

	return &gr, nil
}

func (g *githubAPI) getContribs() (*githubContribs, error) {
	doc, err := g.docGetter(g.config)
	if err != nil {
		return nil, err
	}
	gc := make(githubContribs)
	doc.Find("g > g > rect").Each(func(i int, s *goquery.Selection) {
		rawDate, _ := s.Attr("data-date")
		rawCount, _ := s.Attr("data-count")

		var (
			t         time.Time
			err       error
			timestamp string = "0"
			count     int    = 0
		)

		t, err = time.Parse("2006-01-02", rawDate)
		if err == nil {
			timestamp = strconv.FormatInt(t.Unix(), 10)
		}
		count, _ = strconv.Atoi(rawCount)

		gc[timestamp] = count
	})
	return &gc, nil
}
