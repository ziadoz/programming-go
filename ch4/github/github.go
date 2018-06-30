// Package github provides a Go API for the Github issue tracker.
// See https://developer.github.com/v3/search/#search-issues
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // In Markdown format.
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the Github issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close the resp.Body on all execution paths.
	// (Chapter 5 present 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func (result *IssuesSearchResult) FilterBetweenTimes(after, before time.Time) *IssuesSearchResult {
	filtered := []*Issue{}

	for _, issue := range result.Items {
		if issue.CreatedAt.After(after) && issue.CreatedAt.Before(before) {
			filtered = append(filtered, issue)
		}
	}

	return &IssuesSearchResult{
		TotalCount: len(filtered),
		Items:      filtered,
	}
}

func (result *IssuesSearchResult) SortByCreatedAtDescending() {
	sort.Slice(result.Items, func(i, j int) bool {
		a := result.Items[i]
		b := result.Items[j]
		return a.CreatedAt.After(b.CreatedAt)
	})
}
