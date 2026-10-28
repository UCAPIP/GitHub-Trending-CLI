package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchTrending(ctx context.Context, duration Duration, limit int) ([]Repo, error) {
	since := time.Now().AddDate(0, 0, -duration.Days()).Format("2006-01-02")

	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=created:>%s&sort=stars&order=desc&per_page=%d",
		since, limit,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// if token := os.Getenv("GITHUB_TOKEN"); token != "" {
	// 	req.Header.Set("Authorization", "token "+token)
	// }

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result struct {
		Items []Repo `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return result.Items, nil
}
