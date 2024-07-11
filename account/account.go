package account

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/vanntrong/go-ethkit/secp256k1"
	"github.com/vanntrong/go-ethkit/utils"
	"github.com/vanntrong/go-ethkit/words"
	pbkdf2 "golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

func RandomMnemonic(bitsLength int) (*string, error) {
	// Validate bitsLength
	if !utils.AssertInArrayInt([]int{128, 160, 192, 224, 256}, bitsLength) {
		return nil, fmt.Errorf("invalid bitsLength, must be one of [128, 160, 192, 224, 256]")
	}

	// Calculate sizes
	byteSize := bitsLength / 8
	checksumBits := bitsLength / 32
	totalBits := bitsLength + checksumBits
	wordCount := totalBits / 11

	// Generate entropy
	entropy := utils.RandomBytes(byteSize)

	// Calculate hash of entropy
	hash := sha256.Sum256(entropy)

	// Append hash to entropy
	entropyWithHash := append(entropy, hash[:]...)

	// Create binary string from entropy and hash
	binaryStr := ""
	for _, value := range entropyWithHash {
		binaryStr += fmt.Sprintf("%08b", value)
		if len(binaryStr) >= totalBits {
			binaryStr = binaryStr[:totalBits]
			break
		}
	}

	// Split binary string into words
	mnemonicWords := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		start := i * 11
		binaryWord := binaryStr[start : start+11]
		wordIndex, err := strconv.ParseInt(binaryWord, 2, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing binary string: %v", err)
		}
		mnemonicWords[i] = words.EnglishWords[wordIndex]
	}

	// Join words into mnemonic string
	mnemonic := strings.Join(mnemonicWords, " ")
	return &mnemonic, nil
}

type MnemonicToSeedOptions struct {
	salt string
}

func MnemonicToSeed(mnemonic string, options ...MnemonicToSeedOptions) []byte {
	salt := ""

	if len(options) > 0 {
		salt = options[0].salt
	}

	return pbkdf2.Key([]byte(mnemonic), []byte(salt), 2048, 64, sha512.New)
}

type MnemonicToPrivateKeyOptions struct {
	salt string
}

func MnemonicToPrivateKey(mnemonic string, options ...MnemonicToPrivateKeyOptions) string {
	salt := ""

	if len(options) > 0 {
		salt = options[0].salt
	}

	seed := MnemonicToSeed(mnemonic, MnemonicToSeedOptions{salt: salt})

	mac := hmac.New(sha512.New, []byte("Bitcoin seed"))
	mac.Write(seed)
	hash := mac.Sum(nil)

	privateKey := hash[0:32]

	return hex.EncodeToString(privateKey)
}

func PrivateKeyToAccount(privateKey string) TAccount {
	Q := secp256k1.ScalarMultiplication(privateKey)

	unCompressedPublicKey := fmt.Sprintf("0x04%s%s", Q.X.Text(16), Q.Y.Text(16))

	keccak256 := sha3.NewLegacyKeccak256()
	publicKeyBytes, _ := hex.DecodeString(utils.PadToEven(utils.StripHexPrefix(unCompressedPublicKey)))
	keccak256.Write(publicKeyBytes)
	hash := keccak256.Sum(nil)

	addressHex := hex.EncodeToString(hash[len(hash)-20:])

	address := "0x" + addressHex

	return TAccount{
		address:    address,
		publicKey:  unCompressedPublicKey,
		privateKey: privateKey,
	}
}
