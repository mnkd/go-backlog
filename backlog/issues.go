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
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"projectId"`
	IssueKey    string    `json:"issueKey"`
	KeyID       int64     `json:"keyId"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`

	Assignee struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"assignee"`

	CreatedUser struct {
		ID     int64  `json:"id"`
		UserID string `json:"userId"`
		Name   string `json:"name"`
	} `json:"createdUser"`

	IssueType  IssueType  `json:"issueType"`
	Categories []Category `json:"category"`
}

// IssueRequest represents a request to create/edit an issue.
type IssueRequest struct {
	Summary       string
	Description   string
	ProjectID     int
	PriorityID    int
	CategoryID    int
	IssueTypeID   int
	CategoryIDs   []int
	AssigneeID    *int
	ParentIssueID *int
	State         *string
	StartDate     *string
	DueDate       *string
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

// Delete delete an issue
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
	v.Set("projectId", fmt.Sprintf("%d", r.ProjectID))
	v.Set("issueTypeId", fmt.Sprintf("%d", r.IssueTypeID))
	v.Set("priorityId", fmt.Sprintf("%d", r.PriorityID))
	v.Set("categoryId[]", fmt.Sprintf("%d", r.CategoryID))
	v.Set("summary", r.Summary)
	v.Set("description", r.Description)

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
