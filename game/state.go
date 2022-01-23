package game

import "github.com/bloodboundy/bloodbound-server/player"

type State struct {
	id           string // the game id
	round        uint32 // current round
	daggerIn     uint32 // now, dagger is in which player's hand(player index)
	actionState  string
	playerStates []*player.State
}
