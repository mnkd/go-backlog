package backlog

import (
	"net/url"
	"time"
)

// IssueComment is Backlog issue comment
type IssueComment struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`

	CreatedUser struct {
		ID     int    `json:"id"`
		UserID string `json:"userId"`
		Name   string `json:"name"`
	} `json:"createdUser"`

	ChangeLogs []ChangeLog `json:"changeLog"`
}

// ChangeLog is Backlog issue comment
type ChangeLog struct {
	Field         string `json:"field"`
	NewValue      string `json:"newValue"`
	OriginalValue string `json:"originalValue"`
}

// ListComments lists all issue comments.
//
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-comment-list/
// order: "asc" or "desc"
func (s *IssuesService) ListComments(issueKey string, order string) ([]*IssueComment, *Response, error) {
	u := "issues/" + issueKey + "/comments?order=" + order
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comments := []*IssueComment{}
	resp, err := s.client.Do(req, &comments)
	if err != nil {
		return nil, resp, err
	}
	return comments, resp, nil
}

// CreateComment creates a new comment on the specified issue.
//
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/add-comment/
func (s *IssuesService) CreateComment(issueKey string, comment string) (*IssueComment, *Response, error) {
	u := "issues/" + issueKey + "/comments"
	v := url.Values{}
	v.Set("content", comment)

	req, err := s.client.NewRequest("POST", u, &v)
	if err != nil {
		return nil, nil, err
	}

	issueComment := new(IssueComment)
	resp, err := s.client.Do(req, &issueComment)
	if err != nil {
		return nil, resp, err
	}
	return issueComment, resp, nil
}
