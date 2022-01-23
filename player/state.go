package player

import (
	"github.com/bloodboundy/bloodbound-server/component"
)

type State struct {
	id     string              // the player's id
	index  uint32              // the index of the player
	char   component.Character // character
	tokens []*component.Token  // tokens took
	items  []*component.Item   // items took
}

func NewPlayerState(player *Player, index uint32, char component.Character) *State {
	return &State{
		id:    player.ID(),
		index: index,
		char:  char,
	}
}
