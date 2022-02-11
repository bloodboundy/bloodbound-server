package action

import (
	"context"

	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/bloodboundy/bloodbound-server/ws"
	"github.com/pkg/errors"
)

const AskIntACT actionType = "ask_int"

func init() {
	registerLoader(
		AskIntACT, AskIntActionJSON{},
		func(ctx context.Context, state *game.State, jsi interface{}) (Action, error) {
			jso := jsi.(*AskIntActionJSON)
			return &AskIntAction{
				actionComm: jso.makeActionComm(TargetACT),
				from:       jso.From,
			}, nil
		})
}

type AskIntAction struct {
	actionComm

	from uint32
}

type AskIntActionJSON struct {
	actionJSONComm
	From     uint32 `json:"from"`
	Attacker uint32 `json:"attacker"`
}

func (a *AskIntAction) Dump(ctx context.Context, state *game.State) *AskIntActionJSON {
	return &AskIntActionJSON{
		actionJSONComm: a.makeActionJSONComm(state),
		From:           a.from,
		Attacker:       state.DaggerIn,
	}
}

func (a *AskIntAction) Check(ctx context.Context, state *game.State) error {
	if state.DaggerTarget != a.from {
		return errors.Errorf("not dagger target, current dagger target is %d", state.DaggerTarget)
	}
	return nil
}

func (a *AskIntAction) Apply(ctx context.Context, state *game.State) error {
	state.ResetWantedTo(string(IntACT), string(NoIntACT))
	state.Ints = nil
	ws.PickManager(ctx).BroadCast(a.Dump(ctx, state), state.PlayerIDs()...)
	return nil
}
