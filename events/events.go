package events

type EventCommon struct {
	Platform      string `json:"platform"`
	BotID         string `json:"botId"`
	BroadcasterID string `json:"broadcasterId"`
}

type Message struct {
	EventCommon

	MessageID string `json:"messageId"`
	Message   string `json:"message"`
	ReplyTo   string `json:"replyTo,omitempty"`

	BroadcasterLogin string `json:"broadcasterLogin"`
	BroadcasterName  string `json:"broadcasterName"`

	ChatterID    string `json:"chatterId"`
	ChatterLogin string `json:"chatterLogin"`
	ChatterName  string `json:"chatterName"`
}

type MessageSend struct {
	EventCommon

	Message string `json:"message"`
	ReplyTo string `json:"replyTo,omitempty"`
}
