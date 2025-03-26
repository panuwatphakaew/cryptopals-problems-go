package problems

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func SingleByteXORCipher(input string) (string, error) {
	ciphertext, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	var key byte
	var plaintext []byte
	scoreResult := 0

	for i := 0; i < 256; i++ {
		decrypted := decryptXOR(ciphertext, byte(i))
		score := calculateScore(decrypted)
		if score > scoreResult {
			scoreResult = score
			key = byte(i)
			plaintext = decrypted
		}
	}

	fmt.Println("key is ", string(key))

	return string(plaintext), nil
}

func decryptXOR(ciphertext []byte, key byte) []byte {
	result := make([]byte, len(ciphertext))
	for i := range ciphertext {
		result[i] = ciphertext[i] ^ key
	}
	return result
}

func calculateScore(text []byte) int {
	freq := "etaoin shrdlu"
	score := 0

	for _, b := range text {
		if strings.ContainsRune(freq, rune(b)) {
			score += 1
		}
	}

	return score
}
