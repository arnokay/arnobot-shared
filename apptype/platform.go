package apptype

import (
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/events"
)

type (
	PlatformChatMessageSend = Request[events.MessageSend]
	PlatformStartBot        = Request[data.PlatformToggleBot]
	PlatformStopBot         = Request[data.PlatformToggleBot]
)
