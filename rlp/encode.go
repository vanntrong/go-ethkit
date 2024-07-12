package rlp

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	EmptyString = []byte{0x80}
	EmptyList   = []byte{0xC0}
	TrueValue   = []byte{0x01}
	FalseValue  = []byte{0x80}
)

var ErrNegativeBigInt = errors.New("rlp: cannot encode negative big.Int")

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

	//switch v := input.(type) {
	//case uint8:
	//	fmt.Println("unit8 value", v)
	//case string:
	//	fmt.Println("String:", v)
	//	// Handle string encoding
	//case int:
	//	fmt.Println("Int:", v)
	//	// Handle int encoding
	//case []byte:
	//	fmt.Println("Byte slice:", v)
	//	// Handle byte slice encoding
	//case []interface{}:
	//	fmt.Println("Slice of interfaces:", v)
	//	// Handle list encoding (RLP supports lists)
	//case nil:
	//	fmt.Println("Nil value")
	//// Handle nil value
	//default:
	//	fmt.Println("Unknown type:", reflect.TypeOf(v), "value:", v)
	//	// Handle unknown type or return an error
	//}
}
