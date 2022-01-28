package action

import (
	"context"
	"encoding/json"
	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/pkg/errors"
)

const PassACT actionType = "pass"

func init() {
	registerLoader(
		PassACT,
		func(ctx context.Context, state *game.State, data []byte) (Action, error) {
			paj := &PassActionJSON{}
			if err := json.Unmarshal(data, paj); err != nil {
				return nil, errors.Wrap(err, "unmarshal")
			}
			return &PassAction{
				actionComm: actionComm{t: TargetACT, op: paj.Operator},
				from:       paj.From,
				to:         paj.To,
			}, nil
		})
}

type PassAction struct {
	actionComm
	from uint32
	to   uint32
}

type PassActionJSON struct {
	actionJSONComm
	From uint32 `json:"from"`
	To   uint32 `json:"to"`
}

func (a *PassAction) Check(ctx context.Context, state *game.State) error {
	if err := errIfNotInWanted(a, state); err != nil {
		return err
	}
	if state.DaggerIn != a.from {
		return errors.Errorf("not dagger holder,now dagger is in #%d", state.DaggerIn)
	}
	if int(a.to) > len(state.PlayerStates) {
		return errors.Errorf("target %d invalid, expected [0,%d)", a.to, len(state.PlayerStates))
	}
	return nil
}

func (a *PassAction) Apply(ctx context.Context, state *game.State) error {
	resetWantedTo(state, TargetACT, PassACT)
	state.DaggerIn = a.to
	return nil
}
