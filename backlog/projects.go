package backlog

type ProjectsService service

// Project is Backlog project
type Project struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ProjectKey string `json:"projectKey"`
}

// IssueType is Backlog issueType
type IssueType struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ProjectID    int64  `json:"projectId"`
	Color        string `json:"color"`
	DisplayOrder int64  `json:"displayOrder"`
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
