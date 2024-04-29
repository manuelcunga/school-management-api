package util

import (
	"math/rand"
	"strconv"
	"strings"
)

func GenerateRandomNumber() string {
	// Generate a random number between 1000 and 9999 (inclusive)
	num := rand.Intn(9000) + 1000
	return strconv.Itoa(num)
}

func AddPrefix(phoneNumber, prefix string) string {
	if !strings.HasPrefix(phoneNumber, prefix) {
		phoneNumber = prefix + phoneNumber
	}
	return phoneNumber
}
