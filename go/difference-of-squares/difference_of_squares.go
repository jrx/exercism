package diffsquares

func Square(n int) int {
	return n * n
}

func SquareOfSum(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return Square(res)
}

func SumOfSquares(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += Square(i)
	}
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
