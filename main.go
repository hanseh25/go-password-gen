package passlock

import (
	"crypto/rand"
	"math/big"
)

type PasswordType int

const (
	Random PasswordType = iota
	AlphaNumeric
	Pin
)

func GeneratePassword() string {
	return generate(12, false, false, false, 0)
}

func generate(passwordLength int, includeNumbersFlag bool, includeSymbolsFlag bool, includeUppercaseFlag bool, passwordType int) string {
	var chars string = "abcdefghijklmnopqrstuvwxyz"

	if includeNumbersFlag || passwordType == int(Random) {
		chars += "0123456789"
	}
	if includeSymbolsFlag || passwordType == int(Random) {
		chars += "!@#$%^&*()_+{}[]:;<>,.?/~`"
	}
	if includeUppercaseFlag || passwordType == int(Random) {
		chars += "ABCDEFGHIJKLMNOPQRSTUVXYZ"
	}

	if passwordType == int(Pin) {
		chars = "0123456789"
	}

	password := make([]byte, passwordLength)
	for i := 0; i < passwordLength; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		password[i] = chars[randomIndex.Int64()]
	}

	return string(password)
}
