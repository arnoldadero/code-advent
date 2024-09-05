package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// MineAdventCoin finds the lowest number which when combined with the secret key
// results in an MD5 hash that starts with 'prefixLength' number of zeroes.
func MineAdventCoin(secretKey string, prefixLength int) (int, error) {
	prefix := strings.Repeat("0", prefixLength)
	for i := 1; i < 10000000; i++ {
		// Concatenate secretKey with the current number i
		data := fmt.Sprintf("%s%d", secretKey, i)
		hash := md5.Sum([]byte(data))
		hashStr := hex.EncodeToString(hash[:])

		if strings.HasPrefix(hashStr, prefix) {
			return i, nil
		}
	}
	return 0, errors.New("no valid hash found")
}
