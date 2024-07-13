package rlp

import (
	"github.com/holiman/uint256"
	"math/big"
	"reflect"
)

type IRLPHandler struct {
	encoder encoder
}

type encoder func(reflect.Value) ([]byte, error)

var rlpHandlerMap = map[reflect.Kind]IRLPHandler{
	reflect.Uint8:  {encoder: byteEncoder},
	reflect.Uint32: {encoder: byteEncoder},
	reflect.Uint64: {encoder: byteEncoder},

	reflect.Bool:   {encoder: boolEncoder},
	reflect.String: {encoder: stringEncoder},
}

func isBigInt(value interface{}) bool {
	bigIntType := reflect.TypeOf(&big.Int{})
	bigInt := big.Int{}
	bigIntNoPointerType := reflect.TypeOf(bigInt)
	return reflect.TypeOf(value) == bigIntNoPointerType || reflect.TypeOf(value) == bigIntType
}

func isUint256(value interface{}) bool {
	uint256Type := reflect.TypeOf(&uint256.Int{})
	u := uint256.Int{}
	uint256NoPointerType := reflect.TypeOf(u)

	return reflect.TypeOf(value) == uint256NoPointerType || reflect.TypeOf(value) == uint256Type
}
