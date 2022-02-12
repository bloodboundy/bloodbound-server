package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bloodboundy/bloodbound-server/component"
	"github.com/bloodboundy/bloodbound-server/player"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type State struct {
	ID            string    // related game ID
	Round         uint32    // current Round
	DaggerIn      uint32    // now, dagger is in which player's hand(player index)
	WantedActions []string  // what actions can be applied following
	DoDefaultAt   time.Time // a default action(set by action) will be executed at

	// fields used in Action
	DaggerTarget uint32   // Target
	Ints         []uint32 // AskInt/Int/NoInt
	NoInts       []uint32 // AskInt/Int/NoInt

	PlayerStates []*player.State
}

func NewState(game *Game, players []*player.Player) (*State, error) {
	if len(players) < MIN_PLAYERS {
		return nil, fmt.Errorf("not enough players: expected >=%d, got %d", MIN_PLAYERS, len(players))
	}
	return &State{
		ID:           game.ID(),
		Round:        0,
		DaggerIn:     0,
		PlayerStates: makePlayerStates(players),
	}, nil
}

func makePlayerStates(players []*player.Player) []*player.State {
	index := r.Perm(len(players))
	chars := pickCharacters(len(players))

	var result []*player.State
	for i := 0; i < len(players); i++ {
		result = append(result, player.NewState(players[index[i]], uint32(i), chars[i]))
	}
	return result
}

func pickCharacters(n int) []component.Character {
	var result []component.Character
	if n%2 == 1 {
		result = append(result, component.RandCharN(r, component.SEC_CLAN, 1)...)
		n--
	}
	result = append(result, component.RandCharN(r, component.BLUE_CLAN, n/2)...)
	result = append(result, component.RandCharN(r, component.RED_CLAN, n/2)...)
	return result
}

func (s *State) PlayerIDs() []string {
	var result []string
	for _, ps := range s.PlayerStates {
		result = append(result, ps.ID)
	}
	return result
}

func (s *State) ResetWantedTo(wanted ...string) {
	s.WantedActions = wanted
}
