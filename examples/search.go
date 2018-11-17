package main

import (
	"flag"
	"fmt"
	"os"

	pointers "github.com/f2prateek/go-pointers"
	"github.com/mnkd/go-backlog/backlog"
)

func main() {
	var space string
	flag.StringVar(&space, "s", "", "Backlog space (e.g. abc-inc)")
	flag.Parse()

	if len(space) == 0 {
		fmt.Fprintf(os.Stderr, "Usage:\n\tgo run simple.go -s space_name -k issue_key\n")
		os.Exit(1)
	}

	apiKey := os.Getenv("BACKLOG_API_KEY")
	if len(apiKey) == 0 {
		fmt.Fprintf(os.Stderr, "simple.go needs \"BACKLOG_API_KEY\" environment variable.\n")
		os.Exit(1)
	}

	client := backlog.NewClient(nil, space, apiKey)

	request := backlog.IssueSearchRequest{
		StatusIDs:   []int{2},
		ParentChild: pointers.Int(4),
		Sort:        pointers.String("created"),
		Order:       pointers.String("asc"),
		Count:       pointers.Int(10),
	}

	issues, _, err := client.Issues.Search(request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Results: %d\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("%v: %v (%v)\n", issue.IssueKey, issue.Summary, issue.CreatedUser.Name)
	}
}
