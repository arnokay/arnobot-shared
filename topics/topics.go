package topics

import (
	"strings"

	"arnobot-shared/platform"
)

const (
	Any = "*"
)

const (
	platformPlaceholder      = "{platform}"
	broadcasterIDPlaceholder = "{broadcasterID}"
)

// Auth topics
const (
	AuthProviderTokenGet          = "auth.provider-token.get"
	AuthProviderTokenUpdateTokens = "auth.provider-token.update-tokens"

	AuthSessionTokenValidate = "auth.session-token.validate"
	AuthSessionTokenExchange = "auth.session-token.exchange"
)

type platformBroadcasterTopic string

func (p platformBroadcasterTopic) Build(platform platform.Platform, broadcasterID string) string {
  topic := strings.Replace(string(p), platformPlaceholder, platform.String(), 1)
	topic = strings.Replace(topic, broadcasterIDPlaceholder, broadcasterID, 1)
	return topic
}

const (
	PlatformBroadcasterChatMessageNotify platformBroadcasterTopic = "chat.message.notify.{platform}.{broadcasterID}"
	PlatformBroadcasterChatMessageSend   platformBroadcasterTopic = "chat.message.send.{platform}.{broadcasterID}"
)
