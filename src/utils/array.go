package utils

import (
	"strconv"
)

func IntsToString(numbers []int) []string {
	result := []string{}
	for _, number := range numbers {
		result = append(result, strconv.Itoa(number))
	}
	return result
}

func Permutations[T any](array []T, prefix []T, n int, k int) [][]T {
	permutations := [][]T{}

	if k == 0 {
		permutations = append(permutations, prefix)
		return permutations
	}

	for _, val := range array {
		newPrefix := append(prefix, val)

		permutations = append(permutations, Permutations(array, newPrefix, n, k-1)...)
	}

	return permutations
}
