package events

import "github.com/arnokay/arnobot-shared/platform"

type EventCommon struct {
	Platform      platform.Platform `json:"platform"`
	BotID         string            `json:"botId"`
	BroadcasterID string            `json:"broadcasterId"`
}

type ChatterRole int

const (
	ChatterPleb ChatterRole = iota + 1
	ChatterSub
	ChatterVIP
	ChatterModerator
	ChatterBroadcaster
)

type Message struct {
	EventCommon

	MessageID string `json:"messageId"`
	Message   string `json:"message"`
	ReplyTo   string `json:"replyTo,omitempty"`

	BroadcasterLogin string `json:"broadcasterLogin"`
	BroadcasterName  string `json:"broadcasterName"`

	ChatterID    string      `json:"chatterId"`
	ChatterLogin string      `json:"chatterLogin"`
	ChatterName  string      `json:"chatterName"`
	ChatterRole  ChatterRole `json:"chatterRole"`
}

type MessageSend struct {
	EventCommon

	Message string `json:"message"`
	ReplyTo string `json:"replyTo,omitempty"`
}
