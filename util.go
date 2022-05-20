package gotrees

func Max[T Comparable](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
