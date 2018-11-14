package backlog

// SpaceService is
type SpaceService service

// Priority is Backlog priority in the Backlog space
type Priority struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ListPriorities lists all priorities.
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
