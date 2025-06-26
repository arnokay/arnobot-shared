package platform

import (
	"errors"
	"slices"
)

type Platform string

const (
	Twitch Platform = "twitch"
	Kick   Platform = "kick"
)

var platformValues = []Platform{Twitch, Kick}

func (p Platform) String() string {
	return string(p)
}

func (p Platform) IsEnum() bool {
	return slices.Contains(platformValues, p)
}

func (p Platform) Validate() error {
	if !p.IsEnum() {
		return errors.New("unknown platform")
	}

	return nil
}
