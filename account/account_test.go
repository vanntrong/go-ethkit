package account

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var TestRandomMnemonicData = map[int]int{
	128: 12,
	160: 15,
	192: 18,
	224: 21,
	256: 24,
}

func TestRandomMnemonic(t *testing.T) {
	for k, v := range TestRandomMnemonicData {
		bits, err := RandomMnemonic(k)

		if err != nil || bits == nil {
			t.Fatalf(`Error %v`, err)
		}

		bitsLength := len(strings.Split(*bits, " "))

		if bitsLength != v {
			t.Fatalf(`Expect %d bits with length %d but got length %d`, k, v, bitsLength)
		}
	}
}

var TestMnemonicToPrivateKeyData = map[string]string{
	"champion home gym boss ten wise robust toss art exit minute guard":                "2522a05061799c9fa6f85a32ae011b0faeab854ffa2da1b57f7d6de9faf5408a",
	"flash real trim control submit interest ginger inform crystal pride sword help":   "db637fc3bd9b64c680b3b263fa127a0328deb8f7088a44c8783ed1e36452ce78",
	"brown damage pause evolve quote powder below thought sauce enable plug goose":     "723676346d3816fcbde60350494213206752806f364f1b7de4d6d442ea061d87",
	"gaze spoon father ghost click gift race funny dinner wrong relief hand":           "c29674f6dd63918ef40a770778442e8e7a91d83d996cb48acb24614583a862b0",
	"mention arrest hazard hello firm crouch clutch long tree language girl all":       "bf3aeef13c00bbfba54c10a2bf09c69bde1cefaa24a4944684bbb6b5408d28d1",
	"faith insect nature royal smart merry fish valley shaft author blame glass":       "8da08eb81e31020804d1cf9909e32b0f201bab0c5fd308264f3c3d46066381ff",
	"jungle protect drip proud surround height person hill idle lazy magic body":       "37019479493744d3941d8b6d5983563c7870bab06623e6619ece8e849410ddc7",
	"reveal admit cart during chunk awkward design riot juice sting phrase fall":       "b6183fb3e2e2d79e2f6bb4d4fe80bb2aa11f1a55a9e03e8b23f5cd43e37e7fe7",
	"tail faculty shiver surround oyster pudding flower fabric praise lady trust lock": "a7407a7241ed5a68345caa09e756c466a0af4a28b53b61831afeae020ed42d33",
	"pen rely sock distance bonus joke medal island inner come exchange swear":         "0e5e91189d3f8550411a95c4eeeb3e27f7c1413a636458ac5ea23a6e18800146",
	"wild hurry tank shove lunch satoshi coyote inform math exhaust dad ancient":       "bbe2fecc217c5a95c4c0a4f0f6a0668e539e175eb9258f96ca33fe02aab157eb",
}

func TestMnemonicToPrivateKey(t *testing.T) {
	for k, v := range TestMnemonicToPrivateKeyData {
		privateKey := MnemonicToPrivateKey(k)

		assert.Equal(t, privateKey, v, fmt.Sprintf("Expected %s but got %s", v, privateKey))
	}
}

func TestPrivateKeyToAccount(t *testing.T) {
	account := PrivateKeyToAccount("d29cf91d2300f075c2ca7f7c9be655b7bb8757e2b893e2bf3775338a4e7095e4")

	assert.Equal(t, "0x0456e40b7f474d7703ae7548cd63e78146e0ee19d0e1f2f62c0c39fa33dcca331e151f77baac8b954cd8c27a2a7a8d4e3de30ccf4022393eb3e078236591b8f8b4", account.publicKey, "unexpected public key")
	assert.Equal(t, "0x2362b328cdb20b1e1d9811fde04e0e98ca8f6b66", account.address, "unexpected address")
}
