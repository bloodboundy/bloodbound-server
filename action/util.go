package action

import (
	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/pkg/errors"
)

func resetWantedTo(state *game.State, acts ...actionType) {
	state.WantedActions = nil
	for _, act := range acts {
		state.WantedActions = append(state.WantedActions, string(act))
	}
}

func errIfNotInWanted(action Action, state *game.State) error {
	for _, v := range state.WantedActions {
		if action.Type() == v {
			return nil
		}
	}
	return errors.Errorf("not acceptable, expected: %v", state.WantedActions)
}
