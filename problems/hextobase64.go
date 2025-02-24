package problems

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexRaw []byte) (string, error) {
	dst := make([]byte, hex.DecodedLen(len(hexRaw)))
	n, err := hex.Decode(dst, hexRaw)
	if err != nil {
		return "", err
	}

	result := base64.StdEncoding.EncodeToString(dst[:n])

	return result, nil
}
