package apptype

import (
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/events"
)

type (
	PlatformChatMessageSend = Request[events.MessageSend]
	PlatformStartBot        = Request[data.PlatformBotToggle]
	PlatformStopBot         = Request[data.PlatformBotToggle]
	PlatformBotGetRequest   = Request[data.PlatformBotGet]
	PlatformBotGetResponse  = Response[data.PlatformBot]
)
