package utils

func Prod[Before any, After int](items []Before, mapper func(item Before) After) After {
	result := After(1)
	for _, item := range items {
		result *= mapper(item)
	}
	return result
}
