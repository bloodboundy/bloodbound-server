package player

import (
	"github.com/bloodboundy/bloodbound-server/component"
)

type State struct {
	ID     string              // the player's ID
	index  uint32              // the index of the player
	char   component.Character // character
	tokens []*component.Token  // tokens took
	items  []*component.Item   // items took
}

func NewState(player *Player, index uint32, char component.Character) *State {
	return &State{
		ID:     player.ID(),
		index:  index,
		char:   char,
		tokens: make([]*component.Token, 0, 4),
	}
}

func (s *State) Index() uint32 {
	return s.index
}
