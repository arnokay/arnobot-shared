package events

import "strings"

const (
  defaultSeparator = "."
)

type Event struct {
	Prefix    string
	Event     string
	Command   string
	Separator string
}

func (e *Event) Build() string {
  eventParts := []string{}

  if e.Separator == "" {
    e.Separator = defaultSeparator
  }

  if e.Prefix != "" {
    eventParts = append(eventParts, e.Prefix)
  }

  if e.Event != "" {
    eventParts = append(eventParts, e.Event)
  }

  if e.Command != "" {
    eventParts = append(eventParts, e.Command)
  }

  return strings.Join(eventParts, e.Separator)
}

type Emote struct {
  EmoteID string
  EmoteName string
  Fragments [][]int
}

type SystemChatMessageEvent struct {
  ChannelID string
  ChannelLogin string
  ChannelName string
  ChatterID string
  ChatterLogin string
  ChatterName string
  MessageID string
  MessageText string
  MessageEmotes []Emote
}
