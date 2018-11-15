package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mnkd/go-backlog/backlog"
)

func main() {
	var space, issueKey string
	flag.StringVar(&space, "s", "", "Backlog space (e.g. abc-inc)")
	flag.StringVar(&issueKey, "k", "", "issue key")
	flag.Parse()

	if len(space) == 0 || len(issueKey) == 0 {
		fmt.Fprintf(os.Stderr, "Usage:\n\tgo run simple.go -s space_name -k issue_key\n")
		os.Exit(1)
	}

	apiKey := os.Getenv("BACKLOG_API_KEY")
	if len(apiKey) == 0 {
		fmt.Fprintf(os.Stderr, "simple.go needs \"BACKLOG_API_KEY\" environment variable.\n")
		os.Exit(1)
	}

	client := backlog.NewClient(nil, space, apiKey)
	issue, _, err := client.Issues.Get(issueKey)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v %v %v\n", issue.Summary, issue.IssueKey, issue.Status.Name)

	issueComment, _, err := client.Issues.CreateComment(issue.IssueKey, "Apple")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v %v\n", issueComment.Content, issueComment.CreatedUser.Name)

	comments, _, err := client.Issues.ListComments(issue.IssueKey, "asc")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comments: \n")
	for _, comment := range comments {
		fmt.Printf("  %v %v\n", comment.CreatedUser.Name, comment.Content)
		if len(comment.ChangeLogs) > 0 {
			changelog := comment.ChangeLogs[0]
			fmt.Printf("ChangeLog:  %v %v -> %v\n", changelog.Field, changelog.OriginalValue, changelog.NewValue)
		}
	}
}
