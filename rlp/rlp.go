package rlp

import (
	"fmt"
	"reflect"
)

func getRLPHandler(input interface{}) *IRLPHandler {
	rType := reflect.TypeOf(input)
	arrayHandler = IRLPHandler{encoder: arrayEncoder}

	var handler *IRLPHandler
	//fmt.Println(rType)
	switch rType.Kind() {
	case reflect.Pointer:
		if isBigInt(input) {
			handler = &bigIntHandler
		}
		if isUint256(input) {
			handler = &uintHandler
		}
	case reflect.Struct:
		if isBigInt(input) {
			handler = &bigIntHandler
		}
		if isUint256(input) {
			handler = &uintHandler
		}
	case reflect.Array:
		elmKind := rType.Elem().Kind()
		fmt.Println("type is array of", elmKind)
		handler = &arrayHandler
	case reflect.Slice:
		elmKind := rType.Elem().Kind()
		fmt.Println("type is slice of", elmKind)
		handler = &arrayHandler
	default:
		handler = rlpHandlerMap[rType.Kind()]
	}

	return handler
}
