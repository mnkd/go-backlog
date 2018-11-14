package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mnkd/go-backlog/backlog"
)

func main() {
	var space string
	flag.StringVar(&space, "s", "", "Backlog space (e.g. abc-inc)")
	flag.Parse()

	if len(space) == 0 {
		fmt.Fprintf(os.Stderr, "Usage:\n\tgo run simple.go -s space_name\n")
		os.Exit(1)
	}

	apiKey := os.Getenv("BACKLOG_API_KEY")
	if len(apiKey) == 0 {
		fmt.Fprintf(os.Stderr, "simple.go needs \"BACKLOG_API_KEY\" environment variable.\n")
		os.Exit(1)
	}

	client := backlog.NewClient(nil, space, apiKey)
	projects, _, err := client.Projects.ListAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for i, project := range projects {
		fmt.Printf("%v. %v (%v)\n", i+1, project.Name, project.ProjectKey)

		issueTypes, _, err := client.Projects.ListIssueTypes(project.ProjectKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		for j, issueType := range issueTypes {
			fmt.Printf("%v. %v (%v)\n", j+1, issueType.Name, issueType.Color)
		}
	}
}
