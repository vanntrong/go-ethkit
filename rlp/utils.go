package rlp

import (
	"github.com/vanntrong/go-ethkit/utils"
	"reflect"
)

func getEncodedHeader(size int) ([]byte, int) {
	if size < 56 {
		bytes := make([]byte, size+1)
		bytes[0] = 0x80 + byte(size)
		return bytes, 1
	}

	sizeHex := utils.PadToEven(utils.IntToHex(size))
	sizeBytes := utils.HexToBytes(sizeHex)
	b := len(sizeHex) / 2
	bytes := make([]byte, size+1+len(sizeBytes))
	f := 0xb7 + byte(b)

	bytes[0] = f
	copy(bytes[1:], sizeBytes)

	return bytes, 1 + len(sizeBytes)
}

func getEncodedArrayHeader(size int) ([]byte, int) {
	if size < 56 {
		bytes := make([]byte, size+1)
		bytes[0] = 0xc0 + byte(size)
		return bytes, 1
	}

	sizeHex := utils.PadToEven(utils.IntToHex(size))
	sizeBytes := utils.HexToBytes(sizeHex)
	b := len(sizeHex) / 2
	bytes := make([]byte, size+1+len(sizeBytes))
	f := 0xf7 + byte(b)

	bytes[0] = f
	copy(bytes[1:], sizeBytes)

	return bytes, 1 + len(sizeBytes)
}

func reflectToArrayByte(value reflect.Value) []byte {
	bytes := make([]byte, value.Len())
	for i := 0; i < value.Len(); i++ {
		bytes[i] = byte(value.Index(i).Uint())
	}

	return bytes
}
