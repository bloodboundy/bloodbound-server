package action

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/pkg/errors"
)

const TargetACT actionType = "target"

func init() {
	registerLoader(
		TargetACT, TargetActionJSON{},
		func(ctx context.Context, state *game.State, jsi interface{}) (Action, error) {
			jso := jsi.(*TargetActionJSON)
			return &TargetAction{
				actionComm: jso.makeActionComm(TargetACT),
				to:         jso.To,
			}, nil
		})
}

type TargetAction struct {
	actionComm
	to uint32
}

type TargetActionJSON struct {
	actionJSONComm
	To uint32 `json:"to"`
}

func (a *TargetAction) Dump(ctx context.Context, state *game.State) *TargetActionJSON {
	return &TargetActionJSON{
		actionJSONComm: a.makeActionJSONComm(state),
		To:             a.to,
	}
}

func (a *TargetAction) Check(ctx context.Context, state *game.State) error {
	if state.DaggerIn != a.index {
		return errors.Errorf("not dagger holder,now dagger is in #%d", state.DaggerIn)
	}
	if int(a.to) > len(state.PlayerStates) {
		return errors.Errorf("target %d invalid, expected [0,%d)", a.to, len(state.PlayerStates))
	}
	return nil
}

func (a *TargetAction) Apply(ctx context.Context, state *game.State) error {
	state.ResetWantedTo(string(AskIntACT), string(NoAskIntACT))
	state.DaggerTarget = a.to
	if err := ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...); err != nil {
		return errors.Wrap(err, "BroadCast")
	}
	return nil
}
