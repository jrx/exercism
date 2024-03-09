package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (i int, err error) {

	if n <= 0 {
		return 0, errors.New("input must be greater than zero")
	}

	for i = 0; n != 1; i++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return
}
