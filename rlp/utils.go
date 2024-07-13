package rlp

import "github.com/vanntrong/go-ethkit/utils"

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
