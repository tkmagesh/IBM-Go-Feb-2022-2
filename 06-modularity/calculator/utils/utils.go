package utils

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return !IsEven(n)
}
