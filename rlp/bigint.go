package rlp

import (
	"math/big"
	"reflect"
)

const wordBytes = (32 << (uint64(^big.Word(0)) >> 63)) / 8

func bigIntEncoder(value reflect.Value) ([]byte, error) {
	var bigInt *big.Int
	if value.Kind() == reflect.Struct {
		temp := value.Interface().(big.Int)
		bigInt = &temp // Extract *big.Int from reflect.Value
	} else {
		bigInt = value.Interface().(*big.Int) // Extract *big.Int from reflect.Value
	}
	bitlen := bigInt.BitLen()
	if bitlen <= 64 {
		return byteEncoder(reflect.ValueOf(bigInt.Uint64()))
	}
	// For larger integers, calculate the byte length
	length := ((bitlen + 7) & -8) >> 3
	encoded := make([]byte, 1+length) // 1 byte for the header

	// Set header
	encoded[0] = byte(0x80 + length)

	// Fill in the byte slice
	index := length
	for _, d := range bigInt.Bits() {
		for j := 0; j < wordBytes && index > 0; j++ {
			index--
			encoded[1+index] = byte(d)
			d >>= 8
		}
	}

	return encoded, nil
}

var bigIntHandler = IRLPHandler{encoder: bigIntEncoder}
