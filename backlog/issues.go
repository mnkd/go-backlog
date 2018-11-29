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

	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`

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

	// Backlog の仕様では Version と　Milestone は同じ型になる
	Versions   []Version `json:"versions"`
	Milestones []Version `json:"milestone"`
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
	VersionID     *int
	MilestoneID   *int
	StartDate     *string
	DueDate       *string
}

// IssueSearchRequest represents a request to create/edit an issue.
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-issue-list/
type IssueSearchRequest struct {
	IDs            []int   `url:"id[],omitempty"`             // 課題のID
	ProjectIDs     []int   `url:"projectId[],omitempty"`      // プロジェクトのID
	StatusIDs      []int   `url:"statusId[],omitempty"`       // 状態のID
	PriorityIDs    []int   `url:"priorityId[],omitempty"`     // 優先度のID
	CategoryIDs    []int   `url:"categoryId[],omitempty"`     // カテゴリーのID
	VersionIDs     []int   `url:"versionId[],omitempty"`      // 課題の発生バージョンのID
	MilestoneIDs   []int   `url:"milestoneId[],omitempty"`    // 課題のマイルストーンのID
	IssueTypeIDs   []int   `url:"issueTypeId[],omitempty"`    // 種別のID
	AssigneeIDs    []int   `url:"assigneeId[],omitempty"`     // 担当者のID
	ParentIssueIDs []int   `url:"parentIssueId[],omitempty"`  // 親課題のID
	StartDateSince *string `url:"startDateSince[],omitempty"` // 開始日 (yyyy-MM-dd)
	DueDateSince   *string `url:"dueDateSince,omitempty"`     // 期限日 (yyyy-MM-dd)
	ParentChild    *int    `url:"parentChild,omitempty"`      // 親子課題の条件
	Sort           *string `url:"sort,omitempty"`             // 課題一覧のソートに使用する属性名
	Order          *string `url:"order,omitempty"`            // `asc` または `desc` 指定が無い場合は `desc`
	Keyword        *string `url:"keyword,omitempty"`          // 検索キーワード
	Count          *int    `url:"count,omitempty"`            // 取得上限 (1-100) 指定が無い場合は 20
	offset         *int    `url:"offset,omitempty"`
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

// Search issues.
func (s *IssuesService) Search(request IssueSearchRequest) ([]*Issue, *Response, error) {
	u, _ := addOptions("issues", request)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	issues := []*Issue{}
	resp, err := s.client.Do(req, &issues)
	if err != nil {
		return nil, resp, err
	}
	return issues, resp, nil
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
	if r.VersionID != nil {
		v.Set("versionId[]", fmt.Sprintf("%d", *r.VersionID))
	}
	if r.MilestoneID != nil {
		v.Set("milestoneId[]", fmt.Sprintf("%d", *r.MilestoneID))
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
