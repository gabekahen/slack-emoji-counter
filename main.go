package main

import (
  "fmt"
  "github.com/nlopes/slack"
  "os"
)

func main() {
  TOKEN := os.Getenv("SLACK_TOKEN")
  api := slack.New(TOKEN)
  params := slack.HistoryParameters {"", "0", 100, false, false}

  foo := ReactionEmoji(api, params)
  fmt.Println(foo)
}
