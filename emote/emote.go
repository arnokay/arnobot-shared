package emote

type IEmote interface {
	GetID() string
	GetName() string
	GetOrigin() string
}

type EmoteReader interface {
	Next() bool
	Get() Emote
}

type Emote struct {
	ID         string
	Name       string
	Origin     string
	SizeX      uint16
	SizeY      uint16
	IsAnimated bool
}

type MessageEmote struct {
  Emote Emote
  Positions [][]int
}

type GlobalEmotes struct{}
