package utils

import (
	"math/rand"
)

const (
	Pool = "abcdefghijklmnopqrstuvwxyzABCDEFGQRSTUVWXYZ0123456789"
)

func RandomString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += string(Pool[rand.Intn(len(Pool))])
	}
	return str
}
