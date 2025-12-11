package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Event struct {
	Type    string  `json:"type"`
	Repo    Repo    `json:"repo"`
	Payload Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Commits []Commit `json:"commits"`
	Action  string   `json:"action"`
}

type Commit struct {
	Message string `json:"message"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <username>")
		return
	}

	username := os.Args[1]
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	// fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: Invalid username or API failure")
		return
	}

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commit(s) to %s\n", len(event.Payload.Commits), event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- %s a new issue in %s\n", event.Payload.Action, event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("- %s a pull request in %s\n", event.Payload.Action, event.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)

		}
	}
}
