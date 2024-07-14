package rlp

import (
	"encoding/hex"
	"reflect"
	"strings"
)

var (
	EmptyValue = []byte{0x80}
	EmptyList  = []byte{0xC0}
	TrueValue  = []byte{0x01}
	FalseValue = []byte{0x80}
)

func Encode(input interface{}) (*string, error) {
	rVal := reflect.ValueOf(input)

	handler := getRLPHandler(input)

	if handler == nil {
		return nil, ErrUnSupportInputType
	}

	val, err := handler.encoder(rVal)

	if err != nil {
		return nil, err
	}

	valHex := strings.ToUpper(hex.EncodeToString(val))

	return &valHex, nil
}
