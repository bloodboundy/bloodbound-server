package fsm

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/pkg/errors"
)

const PassACT actionType = "pass"

func init() {
	registerLoader(
		PassACT, PassActionJSON{},
		func(ctx context.Context, state *State, jsi interface{}) (Action, error) {
			jso := jsi.(*PassActionJSON)
			return &PassAction{
				actionComm: jso.makeActionComm(PassACT),
				to:         jso.To,
			}, nil
		})
}

type PassAction struct {
	actionComm
	to uint32
}

type PassActionJSON struct {
	actionJSONComm
	To uint32 `json:"to"`
}

func (a *PassAction) Dump(_ context.Context, state *State) interface{} {
	return &PassActionJSON{
		actionJSONComm: a.makeActionJSONComm(state),
		To:             a.to,
	}
}

func (a *PassAction) Check(_ context.Context, state *State) error {
	if state.DaggerIn != a.index {
		return errors.Errorf("not dagger holder,now dagger is in #%d", state.DaggerIn)
	}
	if int(a.to) > len(state.PlayerStates) {
		return errors.Errorf("target %d invalid, expected [0,%d)", a.to, len(state.PlayerStates))
	}
	return nil
}

func (a *PassAction) Apply(ctx context.Context, state *State) error {
	state.ResetWantedTo(TargetACT, PassACT)
	state.DaggerIn = a.to
	if err := ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...); err != nil {
		return errors.Wrap(err, "BroadCast")
	}
	return nil
}
