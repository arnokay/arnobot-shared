package events

type EventMessage struct {
	Platform string `json:"platform"`
	BotID    string `json:"botId"`

	MessageID string `json:"messageId"`
	Message   string `json:"message"`
	ReplyTo   string `json:"replyTo"`

	BroadcasterID    string `json:"broadcasterId"`
	BroadcasterLogin string `json:"broadcasterLogin"`
	BroadcasterName  string `json:"broadcasterName"`

	ChatterID    string `json:"chatterId"`
	ChatterLogin string `json:"chatterLogin"`
	ChatterName  string `json:"chatterName"`
}
