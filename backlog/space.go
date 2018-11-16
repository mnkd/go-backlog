package backlog

// SpaceService is
type SpaceService service

// Priority is Backlog priority types in the Backlog space
type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Status is Backlog status types in the Backlog space
type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Resolution is Backlog resolution types in the Backlog space
type Resolution struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ListPriorities lists all priorities.
//
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-priority-list/
func (s *SpaceService) ListPriorities() ([]*Priority, *Response, error) {
	u := "priorities"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	priorities := []*Priority{}
	resp, err := s.client.Do(req, &priorities)
	if err != nil {
		return nil, resp, err
	}
	return priorities, resp, nil
}

// ListStatuses lists all statuses.
//
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-status-list/
func (s *SpaceService) ListStatuses() ([]*Status, *Response, error) {
	u := "statuses"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	statuses := []*Status{}
	resp, err := s.client.Do(req, &statuses)
	if err != nil {
		return nil, resp, err
	}
	return statuses, resp, nil
}

// ListResolutions lists all resolutions.
//
// https://developer.nulab-inc.com/ja/docs/backlog/api/2/get-resolution-list/
func (s *SpaceService) ListResolutions() ([]*Resolution, *Response, error) {
	u := "resolutions"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resolutions := []*Resolution{}
	resp, err := s.client.Do(req, &resolutions)
	if err != nil {
		return nil, resp, err
	}
	return resolutions, resp, nil
}
