package action

import (
	"context"
	"encoding/json"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/ws"
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

func (a *PassAction) Dump(ctx context.Context, state *game.State) *PassActionJSON {
	return &PassActionJSON{
		actionJSONComm: actionJSONComm{Type: a.Type(), Operator: a.Operator(), Round: state.Round},
		From:           a.from,
		To:             a.to,
	}
}

func (a *PassAction) Check(ctx context.Context, state *game.State) error {
	if state.DaggerIn != a.from {
		return errors.Errorf("not dagger holder,now dagger is in #%d", state.DaggerIn)
	}
	if int(a.to) > len(state.PlayerStates) {
		return errors.Errorf("target %d invalid, expected [0,%d)", a.to, len(state.PlayerStates))
	}
	return nil
}

func (a *PassAction) Apply(ctx context.Context, state *game.State) error {
	state.ResetWantedTo(string(TargetACT), string(PassACT))
	state.DaggerIn = a.to
	if err := ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...); err != nil {
		return errors.Wrap(err, "BroadCast")
	}
	return nil
}
