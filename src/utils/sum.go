package utils

func Sum[Before any, After int](items []Before, mapper func(item Before) After) After {
	result := After(0)
	for _, item := range items {
		result += mapper(item)
	}
	return result
}
