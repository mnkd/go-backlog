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
		os.Exit(1)
	}

	for i, project := range projects {
		fmt.Printf("%v. %v (%v)\n", i+1, project.Name, project.ProjectKey)

		issueTypes, _, err := client.Projects.ListIssueTypes(project.ProjectKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("IssueTypes: \n")
		for _, issueType := range issueTypes {
			fmt.Printf("  %v (%v)\n", issueType.Name, issueType.Color)
		}

		categories, _, err := client.Projects.ListCategories(project.ProjectKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("Categories: \n")
		for _, category := range categories {
			fmt.Printf("  %v\n", category.Name)
		}

		users, _, err := client.Projects.ListUsers(project.ProjectKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("Users: \n")
		for _, user := range users {
			fmt.Printf("  %v\n", user.Name)
		}
	}
}
