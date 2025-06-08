package platform

type Platform string

const (
	Twitch Platform = "twitch"
)

func (p Platform) String() string {
	return string(p)
}
