package backlog

import (
	"fmt"
	"net/url"
	"time"
)

// IssuesService is
type IssuesService service

// Issue is Backlog issue
type Issue struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectId"`
	IssueKey    string    `json:"issueKey"`
	KeyID       int       `json:"keyId"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`

	Assignee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"assignee"`

	CreatedUser struct {
		ID     int    `json:"id"`
		UserID string `json:"userId"`
		Name   string `json:"name"`
	} `json:"createdUser"`

	IssueType  IssueType  `json:"issueType"`
	Categories []Category `json:"category"`
}

// IssueRequest represents a request to create/edit an issue.
type IssueRequest struct {
	Summary       *string
	Description   *string
	StatusID      *int
	ProjectID     *int
	PriorityID    *int
	CategoryID    *int
	IssueTypeID   *int
	AssigneeID    *int
	ParentIssueID *int
	StartDate     *string
	DueDate       *string
}

// Get an issue.
func (s *IssuesService) Get(issueKey string) (*Issue, *Response, error) {
	u := "issues/" + issueKey
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	issue := new(Issue)
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}
	return issue, resp, nil
}

// Create creates an issue
func (s *IssuesService) Create(request IssueRequest) (*Issue, *Response, error) {
	u := "issues"
	v := request.makeValues()
	req, err := s.client.NewRequest("POST", u, &v)
	if err != nil {
		return nil, nil, err
	}

	issue := new(Issue)
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}
	return issue, resp, nil
}

// Edit an issue
func (s *IssuesService) Edit(issueKey string, request IssueRequest) (*Issue, *Response, error) {
	u := "issues/" + issueKey
	v := request.makeValues()
	req, err := s.client.NewRequest("PATCH", u, &v)
	if err != nil {
		return nil, nil, err
	}

	issue := new(Issue)
	resp, err := s.client.Do(req, &issue)
	if err != nil {
		return nil, resp, err
	}
	return issue, resp, nil
}

// Delete an issue
func (s *IssuesService) Delete(issueKey string) (*Response, error) {
	u := "issues/" + issueKey
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (r IssueRequest) makeValues() url.Values {
	v := url.Values{}
	if r.ProjectID != nil {
		v.Set("projectId", fmt.Sprintf("%d", *r.ProjectID))
	}
	if r.IssueTypeID != nil {
		v.Set("issueTypeId", fmt.Sprintf("%d", *r.IssueTypeID))
	}
	if r.StatusID != nil {
		v.Set("statusId", fmt.Sprintf("%d", *r.StatusID))
	}
	if r.PriorityID != nil {
		v.Set("priorityId", fmt.Sprintf("%d", *r.PriorityID))
	}
	if r.CategoryID != nil {
		v.Set("categoryId[]", fmt.Sprintf("%d", *r.CategoryID))
	}
	if r.Summary != nil {
		v.Set("summary", *r.Summary)
	}
	if r.Description != nil {
		v.Set("description", *r.Description)
	}
	if r.ParentIssueID != nil {
		v.Set("parentIssueId", fmt.Sprintf("%d", *r.ParentIssueID))
	}
	if r.AssigneeID != nil {
		v.Set("assigneeId", fmt.Sprintf("%d", *r.AssigneeID))
	}
	if r.StartDate != nil {
		v.Set("startDate", *r.StartDate)
	}
	if r.DueDate != nil {
		v.Set("dueDate", *r.DueDate)
	}

	return v
}
