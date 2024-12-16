package utils

func Map[Before any, After any](items []Before, mapper func(item Before) After) []After {
	result := make([]After, 0, len(items))
	for _, item := range items {
		result = append(result, mapper(item))
	}
	return result
}
