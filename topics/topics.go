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

type platformBroadcasterTopic = string

const (
	PlatformBroadcasterChatMessageNotify platformBroadcasterTopic = "chat.message.notify.{platform}.{broadcasterID}"
	PlatformBroadcasterChatMessageSend   platformBroadcasterTopic = "chat.message.send.{platform}.{broadcasterID}"
)

type platformBroadcaster struct {
	platform      platform.Platform
	broadcasterID string
}

func PlatformBroadcaster() platformBroadcaster {
	return platformBroadcaster{
		platform:      platform.All,
		broadcasterID: Any,
	}
}

func (t *platformBroadcaster) Platform(platform platform.Platform) *platformBroadcaster {
	t.platform = platform
	return t
}

func (t *platformBroadcaster) BroadcasterID(broadcasterID string) *platformBroadcaster {
	t.broadcasterID = broadcasterID
	return t
}

func (t *platformBroadcaster) Build(topic platformBroadcasterTopic) string {
	topic = strings.Replace(string(topic), platformPlaceholder, t.platform.String(), 1)
	topic = strings.Replace(string(topic), broadcasterIDPlaceholder, t.broadcasterID, 1)
  return topic
}

