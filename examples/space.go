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

	priorities, _, err := client.Space.ListPriorities()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Priorities: \n")
	for _, priority := range priorities {
		fmt.Printf("  %v: %v\n", priority.ID, priority.Name)
	}

	resolutions, _, err := client.Space.ListResolutions()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Resolutions: \n")
	for _, resolution := range resolutions {
		fmt.Printf("  %v: %v\n", resolution.ID, resolution.Name)
	}

	statuses, _, err := client.Space.ListStatuses()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Statuses: \n")
	for _, status := range statuses {
		fmt.Printf("  %v: %v\n", status.ID, status.Name)
	}
}
