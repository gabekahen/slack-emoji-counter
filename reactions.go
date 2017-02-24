package main

import (
  "fmt"
  "github.com/nlopes/slack"
  "sort"
  "os"
)

func main() {
  var TOKEN string = os.Getenv("SLACK_TOKEN")
  api := slack.New(TOKEN)
  params := slack.HistoryParameters {"", "0", 100, false, false}

	channels, err := api.GetChannels(false)
  groups, err := api.GetGroups(true)

  var reactions []string

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
  for _, group := range groups {
    fmt.Println(group.Name)
    history, err := api.GetGroupHistory(group.ID, params)
    if err != nil {
      fmt.Printf("%s\n", err)
    }
    for _, message := range history.Messages {
      if len(message.Reactions) != 0 {
        for _, reaction := range message.Reactions {
          reactions = append(reactions, reaction.Name)
        }
      }
    }
  }
	for _, channel := range channels {
    if ! channel.IsArchived {
      fmt.Println(channel.Name)
      history, err := api.GetChannelHistory(channel.ID, params)
      if err != nil {
    		fmt.Printf("%s\n", err)
    		return
    	}
      for _, message := range history.Messages {
        if len(message.Reactions) != 0 {
          for _, reaction := range message.Reactions {
            reactions = append(reactions, reaction.Name)
          }
        }
      }
    }
  }
  sort.Strings(reactions)
  fmt.Println(reactions)
}
