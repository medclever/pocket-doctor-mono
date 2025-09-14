package utils

import "github.com/matoous/go-nanoid/v2"

const ALPHABET = "abcdefghijklmnopqrstuvwxyz1234567890"

func GenereateUid(size int) string {
	id, _ := gonanoid.Generate(ALPHABET, size)
	return id
}
