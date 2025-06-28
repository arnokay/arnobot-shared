package topics

import (
	"strings"

	"github.com/arnokay/arnobot-shared/platform"
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

type topicBuilder struct {
	original string
	modified string
}

func TopicBuilder(topic string) *topicBuilder {
	return &topicBuilder{
		original: topic,
		modified: topic,
	}
}

func (b *topicBuilder) Platform(platform platform.Platform) *topicBuilder {
	b.modified = strings.Replace(b.modified, platformPlaceholder, platform.String(), 1)
	return b
}

func (b *topicBuilder) BroadcasterID(broadcasterID string) *topicBuilder {
	b.modified = strings.Replace(b.modified, broadcasterIDPlaceholder, broadcasterID, 1)
	return b
}

func (b *topicBuilder) Build() string {
	return b.modified
}

func (b *topicBuilder) Original() string {
	return b.original
}

const (
	PlatformBroadcasterChatMessageNotify = "chat.message.notify.{platform}.{broadcasterID}"
	PlatformBroadcasterChatMessageSend   = "chat.message.send.{platform}.{broadcasterID}"
	PlatformGetBot                       = "bot.{platform}.get"
	PlatformStartBot                     = "bot.{platform}.start"
	PlatformStopBot                      = "bot.{platform}.stop"
)
