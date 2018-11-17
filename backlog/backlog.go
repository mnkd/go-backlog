package backlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

// A Client manages communication with the Backlog API.
type Client struct {
	client  *http.Client // HTTP client
	BaseURL *url.URL
	common  service
	apiKey  string

	// Services
	Space    *SpaceService
	Projects *ProjectsService
	Issues   *IssuesService
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
	c.Space = (*SpaceService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
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
				e.URL = sanitizeURL(url).String()
				return nil, e
			}
		}
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
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
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// An ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Errors   []Error        `json:"errors"` // more detail on individual errors
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d\nErros:\n%v\n",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.moreDetail())
}

func (r *ErrorResponse) moreDetail() string {
	s := ""
	for _, e := range r.Errors {
		s += fmt.Sprintf("  code:%v, message:%v, info:%v\n", e.Code, e.Message, e.MoreInfo)
	}
	return s
}

// An Error reports more details on an individual error in an ErrorResponse.
type Error struct {
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"moreInfo"`
}

// sanitizeURL redacts the apiKey parameter from the URL which may be exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("apiKey")) > 0 {
		params.Set("apiKey", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
// https://github.com/google/go-github/blob/99760a16213d6fdde13f4e477438f876b6c9c6eb/github/github.go#L212-L232
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
