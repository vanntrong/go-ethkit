package rlp

import (
	"math"
	"reflect"
)

var limitStringLengthEncoder = int64(math.Pow(2, 66))

func stringEncoder(value reflect.Value) ([]byte, error) {
	rVal := value.String()
	length := len(rVal)
	if length == 0 {
		return EmptyValue, nil
	}

	if length == 1 {
		rValByte := []byte(rVal)

		return rValByte, nil
	}

	if length <= 55 {
		rValByte, baseIndex := getEncodedHeader(length)
		for i, char := range rVal {
			rValByte[i+baseIndex] = byte(char)
		}

		return rValByte, nil
	}

	if int64(length) >= limitStringLengthEncoder {
		return []byte{}, ErrLimitSizeString
	}

	rValByte, baseIndex := getEncodedHeader(length)
	for i, char := range rVal {
		rValByte[baseIndex+i] = byte(char)
	}

	return rValByte, nil
}
