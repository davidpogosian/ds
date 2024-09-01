package comparators

type Comparator[T any] func(a, b T) int

func ComparatorInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func ComparatorFloat64(a, b float64) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}
