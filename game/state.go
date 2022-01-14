package game

type GameState string

const (
	UNEXIST GameState = "unexist"
	WAITING GameState = "waiting"
	PLAYING GameState = "playing"
	PAUSED  GameState = "paused"
	ENDED   GameState = "ended"
)
