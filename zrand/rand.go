package zrand

import (
	"math/rand"
	"time"
)

var e = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

func Runes(length int) string {
	rand.New(rand.NewSource(time.Now().UnixMilli()))

	b := make([]rune, length)
	for i := range b {
		b[i] = e[rand.Intn(len(e))]
	}
	return string(b)
}
