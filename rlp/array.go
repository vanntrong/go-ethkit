package rlp

import (
	"reflect"
)

func arrayEncoder(value reflect.Value) ([]byte, error) {
	var elementBytes []byte

	// special case: byte array
	if value.Type().Elem().Kind() == reflect.Uint8 {
		bytes := reflectToArrayByte(value)
		return byteArrayEncoder(bytes)
	}

	if value.Len() == 0 {
		return EmptyList, nil
	}

	for i := 0; i < value.Len(); i++ {
		element := value.Index(i)
		handler := getRLPHandler(element.Interface())

		if handler == nil {
			return []byte{}, ErrUnSupportInputType
		}

		bytes, err := handler.encoder(element)
		if err != nil {
			return []byte{}, err
		}
		elementBytes = append(elementBytes, bytes...)
	}
	elementBytesLength := len(elementBytes)
	rValByte, baseIndex := getEncodedArrayHeader(elementBytesLength)
	for i, by := range elementBytes {
		rValByte[baseIndex+i] = by
	}

	return rValByte, nil

}

func byteArrayEncoder(data []byte) ([]byte, error) {
	dataLength := len(data)

	if dataLength == 1 && data[0] <= 0x7f {
		// Single byte, no need for length prefix
		return data, nil
	}
	rValByte, baseIndex := getEncodedHeader(dataLength)
	for i, by := range data {
		rValByte[baseIndex+i] = by
	}

	return rValByte, nil
}

var arrayHandler IRLPHandler
