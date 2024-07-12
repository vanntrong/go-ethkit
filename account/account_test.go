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

var TestPrivateKeyToAccountData = []TAccount{
	{privateKey: "d29cf91d2300f075c2ca7f7c9be655b7bb8757e2b893e2bf3775338a4e7095e4", publicKey: "0x0456e40b7f474d7703ae7548cd63e78146e0ee19d0e1f2f62c0c39fa33dcca331e151f77baac8b954cd8c27a2a7a8d4e3de30ccf4022393eb3e078236591b8f8b4", address: "0x2362b328cdb20b1e1d9811fde04e0e98ca8f6b66"},
	{privateKey: "2afcd9d0c9fc53d442d0a96546c407a0fd9b7a97378378b8a9a133274a0e9df9", publicKey: "0x04754dca4965f4c4de855131135658d06ed2e9d339e5ee49ead532fd8226da9447d4d4e2567f3a0ee502c6ffe453cc4d5036ca25cd92af110189172e9785b69a6f", address: "0x3934b97513379ae42eeb939eab117c812a78be7a"},
	{privateKey: "131c753d17e7b7f7f8fb83bb87c3faeb67192c2b6876b188128cdbfe9b31b5f5", publicKey: "0x045abe4ac08d9d36b5c24078e862968a52323c453d56f02662f9c2021c9e97bcdd6b6eee0671dc5e74cfdb70204c2ac06cd47693ad6e8cb3e26c73fc6a08be2349", address: "0x76412831100fe350b6ff550b9eb5fea5ecc474e9"},
	{privateKey: "e561246d7477b3098a7bb5d9d25a6a936b931f0af2b5216321ee67cbd2112359", publicKey: "0x049a69ee45d8d2b57709c0ccfd1e70bfdfab9c61473ac3720296d11c297366052b3900a472d584e19b7b77ca5bab29a7049a7713d113208624b2d608cb27cbf1c9", address: "0x92738292b89b231cda51beb9e7bc40674a5c2511"},
	{privateKey: "14c0bda2e0c884b48d219a0f3ff772d785182f98f78529f95e2fa2e4ab3b2644", publicKey: "0x047bdd8e53652f973685e1aa59e623b026457d561e150f21722bb866ff52b0676569d40e69e918222da87dd8ea56a3ade96e8e8e37710121e582c225645c73ab06", address: "0x25aaee72155f4c8dfbf25c66c685e6d3b2dbccf4"},
	{privateKey: "f55faadb886465b43d42d45f507d0bdd51fcd76d7a4af0203a84e896d760238f", publicKey: "0x0456b8a8dd200efbc041fa08a4c621867e6124709b3a5aea558941378f468519555a29337a4e369e70fb546e8d3f2e8f1591a883a6c7aa678debf28c592478248d", address: "0xf9ce264ee191bba9f02f851e1e5d95787dd949bd"},
	{privateKey: "418377d33afd7603a366cb488728a777eaba414dce636c1edbf4568929e00114", publicKey: "0x04941d2845c30c3ee665e4f521c1a4b2f54ed33152b8973d2c5d147befcaf6703e9a13dd0285b179048ba75dc164c0f49f7d894bcf5d3af19e7b178b3aa4f9a386", address: "0x9c1508ec60e6db993bd254939f5263268fffdce0"},
	{privateKey: "399f8b669cd4ba6dc0ee4494fb2190f593c2998203bfc43f98dcd2064d968110", publicKey: "0x0409e77480d3902cf283310ac5230b4146aa231e6f81a7121a3027d623c082ec7183d8178360779bc3c5c9f83bbe09c58c4e656646fb55d2c43fe491f1ee0b0c4c", address: "0xcea4ec54d42a2735952eefc0c10c34adbc9c3b74"},
	{privateKey: "a913dcc3f8e7d0341fdd453725b752d5c34545ddf7a113a3eff70c722471ccf0", publicKey: "0x0482ccc1da6f34df58ed5f79b7d13747ce9e2a9e880a41f4909bfb4979a9f2291ec3a73d8fb02b0b132287ae07357b20ec3be0b4a5c4c338efaa8ed191f345b313", address: "0x15c2fd29406b54fcd638c0c0556ff82f067a0128"},
	{privateKey: "bf4034814c46b796a67ea8fb332afc08114a767d87b1ad8e35651ec938718ea3", publicKey: "0x0490e6ded4be1867f5c3479c3bf9213f3e98b5be2f12c8ebd8d83232fb3d8f362ea226b75fd1c749cee444c6ededa68dc97531b26af3be67061a77b6e06ecd19fe", address: "0xf573b6d53d1a11845d2099fd0931efce79d629f5"},
}

func TestPrivateKeyToAccount(t *testing.T) {
	for _, v := range TestPrivateKeyToAccountData {
		account := PrivateKeyToAccount(v.privateKey)

		assert.Equal(t, v.publicKey, account.publicKey, "unexpected public key")
		assert.Equal(t, v.address, account.address, "unexpected address")
	}

}
