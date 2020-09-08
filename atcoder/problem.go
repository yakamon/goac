package atcoder

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

type Problem struct {
	BaseURL   *url.URL
	URL       string
	ContestID string
	TaskID    string
}

func NewProblemFromURL(problemURL string) (*Problem, error) {
	baseURL, err := url.Parse("https://atcoder.jp")
	if err != nil {
		return nil, err
	}
	u, err := baseURL.Parse(problemURL)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(u.Path, "/")
	return &Problem{
		BaseURL:   baseURL,
		URL:       problemURL,
		ContestID: parts[1],
		TaskID:    parts[3],
	}, nil
}

func (p *Problem) ActionPath() string {
	return fmt.Sprintf("/contests/%s/submit", p.ContestID)
}

func (p *Problem) SubmissionURL() string {
	return path.Join(p.BaseURL.String(), fmt.Sprintf("/contests/%s/submissions", p.ContestID))
}
