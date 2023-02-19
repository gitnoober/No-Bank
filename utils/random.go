package utils

import (
	"math/rand"
	"strings"
	"time"
)

const lowercaseAlphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer b/w min and max
func RandomInt(min, max int64) int64 {
	return min * rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(lowercaseAlphabets)
	for i := 0; i < n; i++ {
		c := lowercaseAlphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return int64(rand.Intn(1000))
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, INR}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
