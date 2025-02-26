package problems

import (
	"encoding/hex"
	"errors"
)

func FixedXOR(input1 string, input2 string) (string, error) {
	buffer1, err := hex.DecodeString(input1)
	if err != nil {
		return "", err
	}

	buffer2, err := hex.DecodeString(input2)
	if err != nil {
		return "", err
	}

	if len(buffer1) != len(buffer2) {
		return "", errors.New("invalid inputs")
	}

	result := make([]byte, len(buffer1))

	for i := 0; i < len(buffer1); i++ {
		result[i] = buffer1[i] ^ buffer2[i]
	}

	return hex.EncodeToString(result), nil
}
