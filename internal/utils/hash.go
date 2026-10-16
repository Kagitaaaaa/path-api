package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

func GenerateHashCode() string {
	now := strconv.Itoa(int(time.Now().UnixNano()))

	hash := sha256.Sum256([]byte(now))

	hashString := hex.EncodeToString(hash[:])

	return strings.ToUpper(hashString[:6])
}
