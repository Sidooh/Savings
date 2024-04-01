package utils

import (
	"time"

	"math/rand"
	"strings"
)

const (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

var r *rand.Rand

func Init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

// RandomBool generates a random boolean
func RandomBool() bool {
	return r.Intn(2) >= 1
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(lowerCharSet)

	for i := 0; i < n; i++ {
		c := lowerCharSet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
