package helpers

import (
	"math/rand"
	"strings"
)

func GetFileExt(filename string) string {
	dotIndex := strings.LastIndex(filename, ".")
	nameLength := len(filename)
	if dotIndex == -1 {
		return "jpeg"
	}
	return filename[dotIndex+1:nameLength]
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}