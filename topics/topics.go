package topics

import (
	"strings"

	"arnobot-shared/platform"
)

// Auth topics
const (
	AuthProviderTokenGet          = "auth.provider-token.get"
	AuthProviderTokenUpdateTokens = "auth.provider-token.update-tokens"

	AuthSessionTokenValidate = "auth.session-token.validate"
	AuthSessionTokenExchange = "auth.session-token.exchange"
)

// Platform topics
const (
	PlatformChatMessageSend platformTopic = platformPlaceholder + ".chat.message.send"
)

// Core topics
const (
	CoreChatMessageNotify = "core.chat-message.notify"
)

type platformTopic string

func (t platformTopic) Platform(platform platform.Platform) string {
	return strings.Replace(string(t), platformPlaceholder, platform.String(), 1)
}

const platformPlaceholder = "{platform}"
