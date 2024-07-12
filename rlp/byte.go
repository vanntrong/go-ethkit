package rlp

import "reflect"

func byteEncoder(value reflect.Value) ([]byte, error) {
	rVal := value.Uint()
	if rVal < 0 {
		return []byte{}, ErrNegativeBigInt
	}

	if rVal == 0 {
		return FalseValue, nil
	}

	if rVal <= 127 {
		return []byte{byte(rVal)}, nil
	}

	// For larger values
	var encodedBytes []byte
	for rVal > 0 {
		encodedBytes = append([]byte{byte(rVal & 0xff)}, encodedBytes...)
		rVal >>= 8
	}

	// Prefix the length of the byte array
	length := byte(len(encodedBytes))
	return append([]byte{0x80 + length}, encodedBytes...), nil
}
