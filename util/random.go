package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}

	return sb.String()
}

// RandomOwner generates a random owner name.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random money amount.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency.
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
