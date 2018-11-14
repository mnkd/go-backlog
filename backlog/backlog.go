package backlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// A Client manages communication with the Backlog API.
type Client struct {
	client  *http.Client // HTTP client
	BaseURL *url.URL
	common  service
	apiKey  string

	// Services
	Projects *ProjectsService
}

type service struct {
	client *Client
}

// NewClient returns a new Backlog API client.
func NewClient(httpClient *http.Client, space string, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse("https://" + space + ".backlog.com/api/v2/")
	c := &Client{client: httpClient, BaseURL: baseURL, apiKey: apiKey}
	c.common.client = c
	c.Projects = (*ProjectsService)(&c.common)
	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, params *url.Values) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("apiKey", c.apiKey)
	u.RawQuery = q.Encode()

	// Debug
	// fmt.Println("u.String():", u.String())

	var buf io.ReadWriter
	if params != nil {
		buf = bytes.NewBufferString(params.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	if params != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	}

	return req, nil
}

// Response is a Backlog API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		_, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return response, readErr
		}
		return response, err
	}

	if v != nil {
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}

	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	return errorResponse
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
