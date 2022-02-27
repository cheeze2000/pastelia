package slug

import (
	"math/rand"
	"strings"
	"time"
)

func RandomSlug() string {
	consonants := "bcdfghjklmnpqrstvwxyz"
	vowels := "aeiou"

	rand.Seed(time.Now().UnixNano())

	var slug strings.Builder
	for i := 0; i < 5; i++ {
		slug.WriteByte(consonants[rand.Intn(21)])
		slug.WriteByte(vowels[rand.Intn(5)])
	}

	return slug.String()
}
