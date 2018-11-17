package backlog

import (
	"net/url"
	"reflect"
	"testing"

	pointers "github.com/f2prateek/go-pointers"
)

// https://github.com/google/go-github/blob/99760a16213d6fdde13f4e477438f876b6c9c6eb/github/github_test.go#L761-L778
func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"/?a=b", "/?a=b"},
		{"/?a=b&apiKey=secret", "/?a=b&apiKey=REDACTED"},
		{"/?a=b&apiKey=secret&foo=id", "/?a=b&apiKey=REDACTED&foo=id"},
	}

	for _, tt := range tests {
		inURL, _ := url.Parse(tt.in)
		want, _ := url.Parse(tt.want)

		if got := sanitizeURL(inURL); !reflect.DeepEqual(got, want) {
			t.Errorf("sanitizeURL(%v) returned %v, want %v", tt.in, got, want)
		}
	}
}

func TestAddOptions(t *testing.T) {
	request := IssueSearchRequest{
		IDs:         []int{1, 2, 3},
		ParentChild: pointers.Int(1),
		Sort:        pointers.String("issueType"),
	}

	u, _ := url.Parse("issues")
	v := url.Values{}
	v.Add("id[]", "1")
	v.Add("id[]", "2")
	v.Add("id[]", "3")
	v.Add("parentChild", "1")
	v.Add("sort", "issueType")
	u.RawQuery = v.Encode()
	want := u.String()

	got, _ := addOptions("issues", request)
	if got != want {
		t.Errorf("addOptions(%v) returned %v, want %v", request, got, want)
	}
}
