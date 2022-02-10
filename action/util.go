package action

import (
	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/pkg/errors"
)

func errIfNotInWanted(actionType string, state *game.State) error {
	for _, v := range state.WantedActions {
		if actionType == v {
			return nil
		}
	}
	return errors.Errorf("not acceptable, expected: %v", state.WantedActions)
}
