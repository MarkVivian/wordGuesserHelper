package components

import (
	"math/rand"
	"time"
)

func Shuffle(alphabet []string) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(alphabet), func(i, j int) {
        alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
    })
}