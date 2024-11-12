package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numberStr = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IntRandom(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func StringRandom(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func NumberStringRandom(n int) string {
	var sb strings.Builder
	k := len(numberStr)
	for i := 0; i < n; i++ {
		c := numberStr[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
