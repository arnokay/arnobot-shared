package consts

const (
	LOGGER_APP_KEY     = "app"
	LOGGER_SERVICE_KEY = "service"
)

type mbTopicKey string

const (
	MB_TOPIC_NEW_MESSAGE mbTopicKey = "new-message"
)

type mbQueueKey string

const (
	MB_QUEUE_NEW_MESSAGE mbQueueKey = "new-message"
  KEK mbQueueKey = "kak"
)

type mbRequestKey string

const (
  MB_AUTH_GET_TOKEN mbRequestKey = "auth.get_token"

)
