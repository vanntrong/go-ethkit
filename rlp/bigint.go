package rlp

import (
	"math/big"
	"reflect"
)

// wordBytes is the number of bytes in a big.Word
const wordBytes = (32 << (uint64(^big.Word(0)) >> 63)) / 8

func bigIntEncoder(value reflect.Value) ([]byte, error) {
	var bigInt *big.Int
	// if type is not pointer, convert it to pointer
	if value.Kind() == reflect.Struct {
		temp := value.Interface().(big.Int)
		bigInt = &temp
	} else {
		bigInt = value.Interface().(*big.Int)
	}
	bitLength := bigInt.BitLen()
	if bitLength <= 64 {
		return byteEncoder(reflect.ValueOf(bigInt.Uint64()))
	}

	// For larger integers, calculate the byte length
	length := ((bitLength + 7) & -8) >> 3
	rValBytes, _ := getEncodedHeader(length)
	index := length
	bytesBuf := rValBytes[len(rValBytes)-length:]
	for _, d := range bigInt.Bits() {
		for j := 0; j < wordBytes && index > 0; j++ {
			index--
			bytesBuf[index] = byte(d)
			d >>= 8
		}
	}

	return rValBytes, nil
}

var bigIntHandler = IRLPHandler{encoder: bigIntEncoder}
