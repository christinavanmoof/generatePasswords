package passwords

import "math/rand"

type GeneratePasswords struct {
	Length int
}

func createPassword(length int) string {
	var password string
	// defined character set from which to generate the password
	var possible string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		println(rand.Intn(len(possible)))
	}
	return password
}
