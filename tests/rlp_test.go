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
	{
		val:    veryVeryBigInt,
		output: "B848FFFFFFFFFFFFFFFFF800000000000000001BFFFFFFFFFFFFFFFFC8000000000000000045FFFFFFFFFFFFFFFFC800000000000000001BFFFFFFFFFFFFFFFFF8000000000000000001",
	},

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

	// strings
	{val: "", output: "80"},
	{val: "\x7E", output: "7E"},
	{val: "\x7F", output: "7F"},
	//{val: "\x80", output: "8180"},
	{val: "dog", output: "83646F67"},
	{
		val:    "Lorem ipsum dolor sit amet, consectetur adipisicing eli",
		output: "B74C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C69",
	},
	{
		val:    "Lorem ipsum dolor sit amet, consectetur adipisicing elit",
		output: "B8384C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E7365637465747572206164697069736963696E6720656C6974",
	},
	{
		val:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur mauris magna, suscipit sed vehicula non, iaculis faucibus tortor. Proin suscipit ultricies malesuada. Duis tortor elit, dictum quis tristique eu, ultrices at risus. Morbi a est imperdiet mi ullamcorper aliquet suscipit nec lorem. Aenean quis leo mollis, vulputate elit varius, consequat enim. Nulla ultrices turpis justo, et posuere urna consectetur nec. Proin non convallis metus. Donec tempor ipsum in mauris congue sollicitudin. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Suspendisse convallis sem vel massa faucibus, eget lacinia lacus tempor. Nulla quis ultricies purus. Proin auctor rhoncus nibh condimentum mollis. Aliquam consequat enim at metus luctus, a eleifend purus egestas. Curabitur at nibh metus. Nam bibendum, neque at auctor tristique, lorem libero aliquet arcu, non interdum tellus lectus sit amet eros. Cras rhoncus, metus ac ornare cursus, dolor justo ultrices metus, at ullamcorper volutpat",
		output: "B904004C6F72656D20697073756D20646F6C6F722073697420616D65742C20636F6E73656374657475722061646970697363696E6720656C69742E20437572616269747572206D6175726973206D61676E612C20737573636970697420736564207665686963756C61206E6F6E2C20696163756C697320666175636962757320746F72746F722E2050726F696E20737573636970697420756C74726963696573206D616C6573756164612E204475697320746F72746F7220656C69742C2064696374756D2071756973207472697374697175652065752C20756C7472696365732061742072697375732E204D6F72626920612065737420696D70657264696574206D6920756C6C616D636F7270657220616C6971756574207375736369706974206E6563206C6F72656D2E2041656E65616E2071756973206C656F206D6F6C6C69732C2076756C70757461746520656C6974207661726975732C20636F6E73657175617420656E696D2E204E756C6C6120756C74726963657320747572706973206A7573746F2C20657420706F73756572652075726E6120636F6E7365637465747572206E65632E2050726F696E206E6F6E20636F6E76616C6C6973206D657475732E20446F6E65632074656D706F7220697073756D20696E206D617572697320636F6E67756520736F6C6C696369747564696E2E20566573746962756C756D20616E746520697073756D207072696D697320696E206661756369627573206F726369206C756374757320657420756C74726963657320706F737565726520637562696C69612043757261653B2053757370656E646973736520636F6E76616C6C69732073656D2076656C206D617373612066617563696275732C2065676574206C6163696E6961206C616375732074656D706F722E204E756C6C61207175697320756C747269636965732070757275732E2050726F696E20617563746F722072686F6E637573206E69626820636F6E64696D656E74756D206D6F6C6C69732E20416C697175616D20636F6E73657175617420656E696D206174206D65747573206C75637475732C206120656C656966656E6420707572757320656765737461732E20437572616269747572206174206E696268206D657475732E204E616D20626962656E64756D2C206E6571756520617420617563746F72207472697374697175652C206C6F72656D206C696265726F20616C697175657420617263752C206E6F6E20696E74657264756D2074656C6C7573206C65637475732073697420616D65742065726F732E20437261732072686F6E6375732C206D65747573206163206F726E617265206375727375732C20646F6C6F72206A7573746F20756C747269636573206D657475732C20617420756C6C616D636F7270657220766F6C7574706174",
	},

	// byte arrays
	{val: [0]byte{}, output: "80"},
	{val: [1]byte{0}, output: "00"},
	{val: [1]byte{1}, output: "01"},
	{val: [1]byte{0x7F}, output: "7F"},
	{val: [1]byte{0x80}, output: "8180"},
	{val: [1]byte{0xFF}, output: "81FF"},
	{val: [3]byte{1, 2, 3}, output: "83010203"},
	{val: [57]byte{1, 2, 3}, output: "B839010203000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
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
