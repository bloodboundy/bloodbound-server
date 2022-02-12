package action

import (
	"github.com/pkg/errors"
)

func inUint32s(value uint32, src []uint32) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}
	return false
}

func inStrings(value string, src []string) bool {
	for _, v := range src {
		if v == value {
			return true
		}
	}
	return false
}

func wrapIfErr(err error, msg string) error {
	if err != nil {
		return errors.Wrap(err, msg)
	}
	return nil
}
