package fsm

import (
	"context"

	"github.com/pkg/errors"
)

const NoAskIntACT actionType = "no_ask_int"

type NoAskIntAction struct {
	actionComm
}

type NoAskIntActionJSON struct {
	actionJSONComm
}

func (a *NoAskIntAction) Dump(_ context.Context, state *State) interface{} {
	return &NoAskIntActionJSON{a.makeActionJSONComm(state)}
}

func (a *NoAskIntAction) Check(_ context.Context, state *State) error {
	ps := state.GetPlayerStateByID(a.Operator())
	if state.DaggerTarget == ps.Index() {
		return errors.Errorf("not dagger target, current dagger target is #%v", state.DaggerTarget)
	}
	return nil
}

func (a *NoAskIntAction) Apply(ctx context.Context, state *State) error {
	state.ResetWantedTo(PickTokenACT)
	return nil
}
