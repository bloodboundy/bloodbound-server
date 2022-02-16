package fsm

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type actionType string

const (
	NoAskIntACT actionType = "no_ask_int"

	NoIntACT actionType = "no_int"

	AcceptIntACT actionType = "accept_int"
	RejectIntACT actionType = "reject_int"

	PickTokenACT actionType = "pick_token"

	SkillACT actionType = "skill"
)

// actionLoader load jsi(action json in interface{} type) into Action
// type(jsi)==type(marshalType), marshalType is the arg in registerLoader
type actionLoader func(ctx context.Context, state *State, jsi interface{}) (Action, error)

var loaderMap = map[string]actionLoader{}
var marshalMap = map[string]interface{}{}

func registerLoader(at actionType, marshalType interface{}, loader actionLoader) {
	ats := string(at)
	if _, ok := loaderMap[ats]; ok {
		logrus.Panicf("dup loader: %v", ats)
	}
	loaderMap[ats] = loader
	if _, ok := marshalMap[ats]; ok {
		logrus.Panicf("dup marshalType: %v", ats)
	}
	marshalMap[ats] = marshalType
}

func Load(ctx context.Context, state *State, data []byte) (Action, error) {
	ajc := actionJSONComm{}
	if err := json.Unmarshal(data, &ajc); err != nil {
		return nil, errors.Wrap(err, "Load.Unmarshal")
	}
	if !inStrings(ajc.Type, state.WantedActions) {
		return nil, errors.Errorf("not acceptable action: %v, expected: %v", ajc.Type, state.WantedActions)
	}
	if _, ok := loaderMap[ajc.Type]; !ok {
		return nil, errors.Errorf("action type not found: %v", ajc.Type)
	}

	jsi := reflect.New(reflect.TypeOf(marshalMap[ajc.Type]))
	if err := json.Unmarshal(data, jsi.Interface()); err != nil {
		return nil, errors.Wrap(err, "unmarshal data")
	}
	return loaderMap[ajc.Type](ctx, state, jsi.Interface())
}

type Action interface {
	// Type getter for type
	Type() string

	Check(ctx context.Context, state *State) error

	// Apply process State according to the Action
	Apply(ctx context.Context, state *State) error

	// Dump Action into a json format
	// the interface{} returned should be able to fit json.Marshal()
	Dump(context.Context, *State) interface{}
}

type actionComm struct {
	t     actionType
	op    string
	index uint32
}

func (ac actionComm) Type() string {
	return string(ac.t)
}

func (ac actionComm) Operator() string {
	return ac.op
}

func (ac actionComm) makeActionJSONComm(state *State) actionJSONComm {
	return actionJSONComm{
		Type:     ac.Type(),
		Operator: ac.Operator(),
		Round:    state.Round,
		From:     ac.index,
	}
}

type actionJSONComm struct {
	Type     string `json:"type"`
	Operator string `json:"operator"` // operator player ID
	From     uint32 `json:"from"`     // operator index
	Round    uint32 `json:"round"`
}

func (aj actionJSONComm) makeActionComm(t actionType) actionComm {
	return actionComm{
		t:     t,
		op:    aj.Operator,
		index: aj.From,
	}
}
