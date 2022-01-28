package action

import (
	"context"
	"encoding/json"
	"github.com/bloodboundy/bloodbound-server/game"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type actionType string

const (
	AskIntACT   actionType = "ask_int"
	NoAskIntACT actionType = "no_ask_int"

	IntACT   actionType = "int"
	NoIntACT actionType = "no_int"

	AcceptIntACT actionType = "accept_int"
	RejectIntACT actionType = "reject_int"

	PickTokenACT actionType = "pick_token"

	SkillACT actionType = "skill"
)

type actionLoader func(ctx context.Context, state *game.State, data []byte) (Action, error)

var loaderMap = map[string]actionLoader{}

func registerLoader(at actionType, loader actionLoader) {
	ats := string(at)
	if _, ok := loaderMap[ats]; ok {
		logrus.Panicf("dup loader: %v", ats)
	}
	loaderMap[ats] = loader
}

func Load(ctx context.Context, state *game.State, data []byte) (Action, error) {
	ajc := actionJSONComm{}
	if err := json.Unmarshal(data, &ajc); err != nil {
		return nil, errors.Wrap(err, "Load.Unmarshal")
	}
	if _, ok := loaderMap[ajc.Type]; !ok {
		return nil, errors.Errorf("action type not found: %v", ajc.Type)
	}
	return loaderMap[ajc.Type](ctx, state, data)
}

type Action interface {
	// Type getter for type
	Type() string

	Check(ctx context.Context, state *game.State) error

	// Apply process State according to the Action
	Apply(ctx context.Context, state *game.State) error
}

type actionComm struct {
	t  actionType
	op string
}

func (ac actionComm) Type() string {
	return string(ac.t)
}

func (ac actionComm) Operator() string {
	return ac.op
}

type actionJSONComm struct {
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Round    uint32 `json:"round"`
}
