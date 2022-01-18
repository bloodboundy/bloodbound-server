package net

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// NewRsp constructs a rsp and filling metadata
func NewRsp(rspPtr interface{}) (*Rsp, error) {
	id, err := uuid.New().MarshalText()
	if err != nil {
		return nil, errors.Wrap(err, "uuid.New")
	}

	rsp := &Rsp{
		meta: meta{
			ID: string(id),
		},
	}

	rtrp := reflect.TypeOf(rspPtr)  // reflect type rspPtr => rtrp
	rvrp := reflect.ValueOf(rspPtr) // reflect value rspPtr => rvrp
	if rtrp.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("only accepts ptr")
	}

	rtr := reflect.TypeOf(rsp).Elem() // reflect type *rsp => rtr
	rvr := reflect.ValueOf(rsp)       // reflect value rsp => rvr
	for i := 0; i < rtr.Elem().NumField(); i++ {
		f := rtr.Field(i)
		if f.Type == rtrp {
			rsp.Type = f.Tag.Get("json")
			rvr.Elem().FieldByName(f.Name).Set(rvrp)
		}
	}
	return rsp, nil
}
