package platform

type Platform string

const (
	Twitch Platform = "twitch"
	All    Platform = "*"
)

func (p Platform) String() string {
	return string(p)
}
