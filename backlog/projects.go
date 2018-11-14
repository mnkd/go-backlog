package backlog

import "net/url"

// ProjectsService is
type ProjectsService service

// Project is Backlog project in the Backlog space
type Project struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProjectKey string `json:"projectKey"`
}

// IssueType is Backlog issueType in the Backlog project
type IssueType struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProjectID    int    `json:"projectId"`
	Color        string `json:"color"`
	DisplayOrder int    `json:"displayOrder"`
}

// Category is Backlog category in the Backlog project
type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProjectKey string `json:"projectKey"`
}

// User is Backlog user
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

// ListAll lists all projects.
func (s *ProjectsService) ListAll() ([]*Project, *Response, error) {
	u := "projects"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projects := []*Project{}
	resp, err := s.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}
	return projects, resp, nil
}

// ListIssueTypes lists all issueTypes.
func (s *ProjectsService) ListIssueTypes(projectKey string) ([]*IssueType, *Response, error) {
	u := "projects/" + projectKey + "/issueTypes"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	types := []*IssueType{}
	resp, err := s.client.Do(req, &types)
	if err != nil {
		return nil, resp, err
	}
	return types, resp, nil
}

// ListCategories lists all issueTypes.
func (s *ProjectsService) ListCategories(projectKey string) ([]*Category, *Response, error) {
	u := "projects/" + projectKey + "/categories"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := []*Category{}
	resp, err := s.client.Do(req, &categories)
	if err != nil {
		return nil, resp, err
	}
	return categories, resp, nil
}

// ListProjectUsers lists all users in the project.
func (s *ProjectsService) ListProjectUsers(projectKey string) ([]*User, *Response, error) {
	u := "projects/" + projectKey + "/users"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := []*User{}
	resp, err := s.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}
	return users, resp, nil
}

// CreateCategory creates a new category in the project.
func (s *ProjectsService) CreateCategory(projectKey string, categoryName string) (*Category, *Response, error) {
	u := "projects/" + projectKey + "/categories"

	v := url.Values{}
	v.Set("name", categoryName)

	req, err := s.client.NewRequest("POST", u, &v)
	if err != nil {
		return nil, nil, err
	}

	category := new(Category)
	resp, err := s.client.Do(req, &category)
	if err != nil {
		return nil, resp, err
	}
	return category, resp, nil
}

// DeleteCategory deletes a category in the project.
func (s *ProjectsService) DeleteCategory(projectKey string, categoryID string) (*Response, error) {
	u := "projects/" + projectKey + "/categories/" + categoryID
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

// CreateIssueType creates a new category in the project.
func (s *ProjectsService) CreateIssueType(projectKey string, name string, color string) (*IssueType, *Response, error) {
	u := "projects/" + projectKey + "/issueTypes"

	v := url.Values{}
	v.Set("name", name)
	v.Set("color", color)

	req, err := s.client.NewRequest("POST", u, &v)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(IssueType)
	resp, err := s.client.Do(req, &issueType)
	if err != nil {
		return nil, resp, err
	}
	return issueType, resp, nil
}

// DeleteIssueType deletes a category in the project.
//
// substituteIssueTypeID: 付け替え先の種別 ID。Backlog の仕様上、最低 1 個の種別を残す必要あり。
func (s *ProjectsService) DeleteIssueType(projectKey string, issueTypeID string, substituteIssueTypeID string) (*Response, error) {
	u := "projects/" + projectKey + "/issueTypes/" + issueTypeID

	// substituteIssueTypeId (必須) 数値 紐づく課題を付け替える先の種別のID
	v := url.Values{}
	v.Set("substituteIssueTypeId", substituteIssueTypeID)

	req, err := s.client.NewRequest("DELETE", u, &v)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
