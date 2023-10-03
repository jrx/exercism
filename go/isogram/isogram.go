package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)

	letters := []rune(word)
	presence := make(map[rune]int)

	for i := 0; i < len(letters); i++ {
		presence[letters[i]] = 0
	}

	return len(presence) == len(letters)
}
