package rlp

import (
	"encoding/hex"
	"fmt"
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
	rType := reflect.TypeOf(input)

	var handler IRLPHandler
	fmt.Println(rType)
	switch rType.Kind() {
	case reflect.Pointer:
		if isBigInt(input) {
			handler = bigIntHandler
		}
		if isUint256(input) {
			//	TODO add handler for uint256
		}
	case reflect.Struct:
		if isBigInt(input) {
			handler = bigIntHandler
		}
		if isUint256(input) {
			//	TODO add handler for uint256
		}
	default:
		handler = rlpHandlerMap[rType.Kind()]
	}
	val, err := handler.encoder(rVal)

	if err != nil {
		return nil, err
	}

	valHex := strings.ToUpper(hex.EncodeToString(val))

	return &valHex, nil
}
