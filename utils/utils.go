package utils

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func DelSlice[T any](slice []T, s int) []T {
	dest := make([]T, len(slice))
	copy(dest, slice)
	return append(dest[:s], dest[s+1:]...)
}
