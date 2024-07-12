package tests

import (
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/vanntrong/go-ethkit/rlp"
	"github.com/vanntrong/go-ethkit/utils"
	"math/big"
	"testing"
)

type rlpTest struct {
	val           interface{}
	output, error string
}

var (
	veryBigInt = new(big.Int).Add(
		new(big.Int).Lsh(big.NewInt(0xFFFFFFFFFFFFFF), 16),
		big.NewInt(0xFFFF),
	)
	veryVeryBigInt = new(big.Int).Exp(veryBigInt, big.NewInt(8), nil)
)

var rlpTestCases = []rlpTest{
	// booleans
	{val: true, output: "01"},
	{val: false, output: "80"},

	// integers
	{val: uint32(0), output: "80"},
	{val: uint32(127), output: "7F"},
	{val: uint32(128), output: "8180"},
	{val: uint32(256), output: "820100"},
	{val: uint32(1024), output: "820400"},
	{val: uint32(0xFFFFFF), output: "83FFFFFF"},
	{val: uint32(0xFFFFFFFF), output: "84FFFFFFFF"},
	{val: uint64(0xFFFFFFFF), output: "84FFFFFFFF"},
	{val: uint64(0xFFFFFFFFFF), output: "85FFFFFFFFFF"},
	{val: uint64(0xFFFFFFFFFFFF), output: "86FFFFFFFFFFFF"},
	{val: uint64(0xFFFFFFFFFFFFFF), output: "87FFFFFFFFFFFFFF"},
	{val: uint64(0xFFFFFFFFFFFFFFFF), output: "88FFFFFFFFFFFFFFFF"},

	// big integers (should match uint for small values)
	{val: big.NewInt(0), output: "80"},
	{val: big.NewInt(1), output: "01"},
	{val: big.NewInt(127), output: "7F"},
	{val: big.NewInt(128), output: "8180"},
	{val: big.NewInt(256), output: "820100"},
	{val: big.NewInt(1024), output: "820400"},
	{val: big.NewInt(0xFFFFFF), output: "83FFFFFF"},
	{val: big.NewInt(0xFFFFFFFF), output: "84FFFFFFFF"},
	{val: big.NewInt(0xFFFFFFFFFF), output: "85FFFFFFFFFF"},
	{val: big.NewInt(0xFFFFFFFFFFFF), output: "86FFFFFFFFFFFF"},
	{val: big.NewInt(0xFFFFFFFFFFFFFF), output: "87FFFFFFFFFFFFFF"},
	{
		val:    new(big.Int).SetBytes(utils.HexToBytes("102030405060708090A0B0C0D0E0F2")),
		output: "8F102030405060708090A0B0C0D0E0F2",
	},
	{
		val:    new(big.Int).SetBytes(utils.HexToBytes("0100020003000400050006000700080009000A000B000C000D000E01")),
		output: "9C0100020003000400050006000700080009000A000B000C000D000E01",
	},
	{
		val:    new(big.Int).SetBytes(utils.HexToBytes("010000000000000000000000000000000000000000000000000000000000000000")),
		output: "A1010000000000000000000000000000000000000000000000000000000000000000",
	},
	{
		val:    veryBigInt,
		output: "89FFFFFFFFFFFFFFFFFF",
	},
	// TODO: apply string for bigint
	//{
	//	val:    veryVeryBigInt,
	//	output: "B848FFFFFFFFFFFFFFFFF800000000000000001BFFFFFFFFFFFFFFFFC8000000000000000045FFFFFFFFFFFFFFFFC800000000000000001BFFFFFFFFFFFFFFFFF8000000000000000001",
	//},

	// non-pointer big.Int
	{val: *big.NewInt(0), output: "80"},
	{val: *big.NewInt(0xFFFFFF), output: "83FFFFFF"},

	// uint256
	{val: uint256.NewInt(0), output: "80"},
	{val: uint256.NewInt(1), output: "01"},
	{val: uint256.NewInt(127), output: "7F"},
	{val: uint256.NewInt(128), output: "8180"},
	{val: uint256.NewInt(256), output: "820100"},
	{val: uint256.NewInt(1024), output: "820400"},
	{val: uint256.NewInt(0xFFFFFF), output: "83FFFFFF"},
	{val: uint256.NewInt(0xFFFFFFFF), output: "84FFFFFFFF"},
	{val: uint256.NewInt(0xFFFFFFFFFF), output: "85FFFFFFFFFF"},
	{val: uint256.NewInt(0xFFFFFFFFFFFF), output: "86FFFFFFFFFFFF"},
	{val: uint256.NewInt(0xFFFFFFFFFFFFFF), output: "87FFFFFFFFFFFFFF"},
	{
		val:    new(uint256.Int).SetBytes(utils.HexToBytes("102030405060708090A0B0C0D0E0F2")),
		output: "8F102030405060708090A0B0C0D0E0F2",
	},
	{
		val:    new(uint256.Int).SetBytes(utils.HexToBytes("0100020003000400050006000700080009000A000B000C000D000E01")),
		output: "9C0100020003000400050006000700080009000A000B000C000D000E01",
	},
}

func TestRLPEncode(t *testing.T) {
	for _, v := range rlpTestCases {
		value, err := rlp.Encode(v.val)
		if err != nil {
			t.Fatalf("Error %s", err)
		}

		assert.Equal(t, v.output, *value)
	}

	//rlp.Encode("hello")
	//rlp.Encode(42)
	//rlp.Encode([]byte{0x01, 0x02, 0x03})
	//rlp.Encode([]interface{}{"hello", 42, []byte{0x01}})
	//rlp.Encode(nil)
	//rlp.Encode(byte(0x00))
	//rlp.Encode(byte(0x7f))
}
