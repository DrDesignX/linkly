package utils

import (
	"math/rand"
)

const (
	Pool = "abcdefghijklmnopqrstuvwxyzABCDEFGQRSTUVWXYZ0123456789"
)

func RandomString(i int) string {
	var str string
	for i < len(Pool) {
		str += string(Pool[rand.Intn(len(Pool))+1])
		i += 1
	}
	return str
}
