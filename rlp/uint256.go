package rlp

import (
	"github.com/holiman/uint256"
	"reflect"
)

func uint256Encoder(value reflect.Value) ([]byte, error) {
	var uInt *uint256.Int

	if value.Kind() == reflect.Struct {
		temp := value.Interface().(uint256.Int)
		uInt = &temp
	} else {
		uInt = value.Interface().(*uint256.Int)
	}

	return bigIntEncoder(reflect.ValueOf(uInt.ToBig()))
}

var uintHandler = IRLPHandler{encoder: uint256Encoder}
