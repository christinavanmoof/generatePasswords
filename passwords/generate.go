package passwords

import (
	"math/rand"
	"time"
)

type GeneratePasswords struct{}

// Generate returns a random password
func (GeneratePasswords) Generate() string {
	charSet := "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"
	pass := RandomStringGenerator(charSet, 8)
	return pass
}

// RandomStringGenerator generates a random string of length n
func RandomStringGenerator(charSet string, codeLength int32) string {
	code := ""
	charSetLength := int32(len(charSet))
	for i := int32(0); i < codeLength; i++ {
		index := randomNumber(0, charSetLength)
		code += string(charSet[index])
	}

	return code
}

// GeneratePassword returns a random password
func randomNumber(min, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return min + int32(rand.Intn(int(max-min)))
}
