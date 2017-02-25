package main

import (
  "fmt"
  "github.com/nlopes/slack"
  "os"
  "log"
)

func main() {
  TOKEN := os.Getenv("SLACK_TOKEN")
  api := slack.New(TOKEN)
  params := slack.HistoryParameters {"", "0", 100, false, false}

  emojilist, err := api.GetEmoji()

  // TODO compile a regular expression from emojilist, and run through
  // regexp.FindAllString.
  fmt.Println(ChannelMessageEmoji(api, params, emojilist)
}
