# GitHub Activity CLI

A simple command-line tool that fetches and displays recent activity for any GitHub user.

## Features

- Fetches recent public events from GitHub's API
- Displays activity in a readable format
- Supports multiple event types:
  - Push events (commits)
  - Issues (opened/closed)
  - Pull requests
  - Stars (watch events)
  - And other GitHub events

## Prerequisites

- Go 1.16 or higher installed on your system

## Installation

1. Clone this repository or download the source code
2. Navigate to the project directory

## Usage

You can either run the program directly or build it first.

### Option 1: Run directly (no build required)

```bash
go run main.go <username>
```

### Option 2: Build and run

Initialize the Go module and build:

```bash
go mod init github-activity
go build -o github-activity
```

Then run the executable:

```bash
./github-activity <username>
```

### Example

```bash
./github-activity torvalds
```

### Sample Output

```
- Pushed 3 commit(s) to torvalds/linux
- Opened a new issue in torvalds/linux
- Starred someuser/awesome-repo
- Opened a pull request in torvalds/subsurface
```

## How It Works

The tool uses GitHub's public Events API endpoint (`https://api.github.com/users/{username}/events`) to fetch the most recent public activity for a specified user. It then parses the JSON response and formats the output in a human-readable way.

## Error Handling

The program handles several error cases:
- Missing username argument
- Invalid username or API failures
- Network errors
- JSON parsing errors

## API Rate Limiting

GitHub's API has rate limits for unauthenticated requests (60 requests per hour). For higher rate limits, you can modify the code to include a GitHub personal access token in the request headers.

## Project URL

This project is part of the roadmap.sh backend projects: https://roadmap.sh/projects/github-user-activity

## License

This is free and unencumbered software released into the public domain.