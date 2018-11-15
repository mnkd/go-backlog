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
	_, _, err := client.Issues.Get("INVALID-ISSUE-KEY")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)

		if errRes, ok := err.(*backlog.ErrorResponse); ok {
			fmt.Fprintf(os.Stderr, "Errors:\n")
			for _, e := range errRes.Errors {
				fmt.Fprintf(os.Stderr, "  code:%v, message:%v, info:%v\n", e.Code, e.Message, e.MoreInfo)
			}
		}

		os.Exit(1)
	}
}
