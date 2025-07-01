package data

type ChatterRole int

const (
	ChatterPleb ChatterRole = iota + 1
	ChatterSub
	ChatterVIP
	ChatterModerator
	ChatterBroadcaster
)

