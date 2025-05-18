// Package for message broker topics.
//
// Naming convention for app-to-app topics:
// <service>.<resource>.<action>
// e.g. auth.token.get
//
// Naming conventioon for external-to-app topics:
// <platform>.<domain>.<resource>.<action>
// e.g. twitch.chat.message.notify
package topics

// Auth topics
const (
	AuthProviderTokenGet     = "auth.provider-token.get"
	AuthSessionTokenValidate = "auth.session-token.validate"
	AuthSessionTokenExchange = "auth.session-token.exchange"
)

// Twitch chat topics
const (
	TwitchChatMessageNotify = "twitch.chat.message.notify"
	TwitchChatMessageSend   = "twitch.chat.message.send"
)
