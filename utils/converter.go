package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

func HexToBigInt(hex string) (*big.Int, error) {
	i := big.NewInt(0)
	if _, ok := i.SetString(hex, 16); !ok {
		return nil, fmt.Errorf("invalid number")
	}
	return i, nil
}

func PadToEven(value string) string {
	if len(value)%2 != 0 {
		value = "0" + value
	}
	return value
}

func StripHexPrefix(value string) string {
	if !isHexPrefixed(value) {
		return value
	}
	return value[2:]
}

func PadToEvenPublicKey(publicKey string) string {
	if len(publicKey)%2 == 0 {
		return publicKey
	}
	return publicKey[:4] + "0" + publicKey[4:]
}

func HexToBytes(val string) []byte {
	bytes, err := hex.DecodeString(val)

	if err != nil {
		return []byte{}
	}

	return bytes
}

func isHexPrefixed(value string) bool {
	if len(value) < 2 {
		return false
	}
	return value[:2] == "0x"
}
