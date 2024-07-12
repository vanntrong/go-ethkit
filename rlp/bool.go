package rlp

import (
	"errors"
	"reflect"
)

var ErrorNotBoolean = errors.New("value is not boolean")

func boolEncoder(value reflect.Value) ([]byte, error) {
	rVal := value.Bool()

	if rVal == true {
		return TrueValue, nil
	}

	if rVal == false {
		return FalseValue, nil
	}

	return []byte{}, ErrorNotBoolean
}
