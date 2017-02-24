package main

import (
  "fmt"
  "github.com/nlopes/slack"
  "log"
)

// Return a string slice containing all individual emoji used
// for reactions within all public channels.
func GroupReactionEmoji(api *slack.Client, params slack.HistoryParameters) []string {
  var reactions []string
  groups, err := api.GetGroups(true)

	if err != nil {
		log.Fatal(err)
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
  return reactions
}

// Return a string slice containing all individual emoji used
// for reactions within visible private groups.
func ChannelReactionEmoji(api *slack.Client, params slack.HistoryParameters) []string {
  var reactions []string
  channels, err := api.GetChannels(false)

	if err != nil {
		log.Fatal(err)
	}
	for _, channel := range channels {
    if ! channel.IsArchived {
      fmt.Println(channel.Name)
      history, err := api.GetChannelHistory(channel.ID, params)
      if err != nil {
    		log.Fatal(err)
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
  return reactions
}
