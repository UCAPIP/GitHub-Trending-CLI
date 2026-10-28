package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	durationStr := flag.String("duration", "week", "time range: day, week, month, year")
	limit := flag.Int("limit", 10, "number of repositories (max 100)")
	flag.Parse()

	duration := Duration(*durationStr)
	if !duration.IsValid() {
		fmt.Fprintf(os.Stderr, "Error: invalid duration %q\n", *durationStr)
		os.Exit(1)
	}
	if *limit < 1 || *limit > 100 {
		fmt.Fprintf(os.Stderr, "Error: limit must be between 1 and 100\n")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	repos, err := FetchTrending(ctx, duration, *limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	PrintTable(repos)
}
