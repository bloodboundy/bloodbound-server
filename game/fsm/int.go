package fsm

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/pkg/errors"
)

const IntACT actionType = "int"

func init() {
	registerLoader(IntACT, IntActionJSON{},
		func(ctx context.Context, state *State, jsi interface{}) (Action, error) {
			jso := jsi.(*IntActionJSON)
			return &IntAction{
				actionComm: jso.makeActionComm(IntACT),
				to:         jso.To,
			}, nil
		})
}

type IntAction struct {
	actionComm
	to uint32
}

type IntActionJSON struct {
	actionJSONComm
	To uint32 `json:"to"`
}

func (a *IntAction) Dump(ctx context.Context, state *State) *IntActionJSON {
	return &IntActionJSON{
		actionJSONComm: a.makeActionJSONComm(state),
		To:             a.to,
	}
}

func (a *IntAction) Check(ctx context.Context, state *State) error {
	if state.DaggerTarget == a.index {
		return errors.Errorf("can not intervene yourself")
	}
	if state.DaggerTarget != a.to {
		return errors.Errorf("#%d is not attack target, expected #%d", a.to, state.DaggerTarget)
	}
	if inUint32s(a.index, state.Ints) {
		return errors.Errorf("already intervened")
	}
	return nil
}

func (a *IntAction) Apply(ctx context.Context, state *State) error {
	state.Ints = append(state.Ints, a.index)
	return wrapIfErr(ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...), "BroadCast")
}
