package random_string

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateRandomStrings(count int, length int) []string {
	strings := make([]string, count)
	for i := range strings {
		strings[i] = randString(length)
	}
	return strings
}

func RandomNumber(low int, high int) int {
	count := high - low
	randomInt := rand.Int()
	return (randomInt % count) + low
}
