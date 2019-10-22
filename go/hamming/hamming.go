package hamming

import (
	"errors"
)

// Distance func calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {

	strandA := []rune(a)
	strandB := []rune(b)
	if len(strandA) != len(strandB) {
		return 0, errors.New("length of the sequences differ")
	}

	dis := 0
	for i := 0; i < len(strandA); i++ {
		if strandA[i] != strandB[i] {
			dis++
		}
	}

	return dis, nil
}
