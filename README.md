# LeetCode Question Fetcher
This is the LeetCode question fetcher that written by Golang

This repository uses GraphQL to fetch the LeetCode Question and the data stores in PostgreSQL

## Linux Environment Deployment

### Prerequisite
1. Database [PostgreSQL]
2. Create a `cert` directory
3. Populate the database query

### Deployment
- Clone the repository<br/>
`git clone git@github.com:WeeHong/leetcode-question-fetcher.git`
- Build the project<br/>
`GOOS=linux GOARCH=amd64 go build -o ./LeetCode-Tracker main.go`
- SystemD setup
```
[Unit]
Description=LeetCode Fetcher

[Service]
Type=simple

Restart=on-failure
RestartSec=10

WorkingDirectory=/root/workspace/go/src/github.com/weehong/leetcode-question-fetcher
ExecStart=/root/workspace/go/src/github.com/weehong/leetcode-question-fetcher/LeetCode-Tracker -ssl

[Install]
WantedBy=multi-user.target
```
- SystemD Timer
```
[Unit]
Description=Schedule the LeetCode Fetcher
Requires=/etc/systemd/system/leetcode-fetcher.service

[Timer]
Unit=/etc/systemd/system/leetcode-fetcher.service
OnCalendar=*-*-* 00:00:00

[Install]
WantedBy=timers.target
```
