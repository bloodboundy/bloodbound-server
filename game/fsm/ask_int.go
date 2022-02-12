package fsm

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/pkg/errors"
)

const AskIntACT actionType = "ask_int"

func init() {
	registerLoader(
		AskIntACT, AskIntActionJSON{},
		func(ctx context.Context, state *State, jsi interface{}) (Action, error) {
			jso := jsi.(*AskIntActionJSON)
			return &AskIntAction{jso.makeActionComm(TargetACT)}, nil
		})
}

type AskIntAction struct {
	actionComm
}

type AskIntActionJSON struct {
	actionJSONComm
	Attacker uint32 `json:"attacker"`
}

func (a *AskIntAction) Dump(ctx context.Context, state *State) *AskIntActionJSON {
	return &AskIntActionJSON{
		actionJSONComm: a.makeActionJSONComm(state),
		Attacker:       state.DaggerIn,
	}
}

func (a *AskIntAction) Check(ctx context.Context, state *State) error {
	if state.DaggerTarget != a.index {
		return errors.Errorf("not dagger target, current dagger target is %d", state.DaggerTarget)
	}
	return nil
}

func (a *AskIntAction) Apply(ctx context.Context, state *State) error {
	state.ResetWantedTo(string(IntACT), string(NoIntACT))
	state.Ints = nil
	ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...)
	return nil
}
