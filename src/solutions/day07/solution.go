package day07

import (
	"aoc/src/utils"
	"math"
	"strconv"
	"strings"
)

func GetPermutations(numbers []int, operators []string) [][]string {
	base := len(operators)
	numOperators := len(numbers) - 1
	numCombos := int(math.Pow(float64(len(operators)), float64(len(numbers)-1)))

	permutations := [][]string{}
	for i := 0; i < numCombos; i++ {
		binStr := strconv.FormatInt(int64(i), base)
		operations := strings.Repeat("0", numOperators-len(binStr)) + binStr

		permutations = append(permutations, strings.Split(operations, ""))
	}

	return permutations
}

func compute(numbers []int, permutations [][]string, expected int) bool {
	for _, operations := range permutations {
		prev := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			current := numbers[i+1]
			if operations[i] == "0" {
				prev = prev + current
			} else if operations[i] == "1" {
				prev = prev * current
			} else if operations[i] == "2" {
				s1 := strconv.Itoa(prev)
				s2 := strconv.Itoa(current)
				prev, _ = strconv.Atoi(s1 + s2)
			}

			if prev > expected {
				break
			}
		}

		if expected == prev {
			return true
		}
	}

	return false
}

func SolveProblem1() (string, error) {
	lines := strings.Split(utils.GetInput(7), "\n")
	operators := []string{"+", "*"}

	expectations := []int{}
	numbers := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")

		expectation, _ := strconv.Atoi(parts[0])
		expectations = append(expectations, expectation)

		values := []int{}
		for _, number := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(number)
			values = append(values, n)
		}
		numbers = append(numbers, values)
	}

	sum := 0
	for i := 0; i < len(expectations); i++ {
		expectation := expectations[i]
		values := numbers[i]

		permutations := GetPermutations(values, operators)
		if compute(values, permutations, expectation) {
			sum += expectation
		}
	}

	return strconv.Itoa(sum), nil
}

func SolveProblem2() (string, error) {
	lines := strings.Split(utils.GetInput(7), "\n")
	operators := []string{"+", "*", "||"}

	expectations := []int{}
	numbers := [][]int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")

		expectation, _ := strconv.Atoi(parts[0])
		expectations = append(expectations, expectation)

		values := []int{}
		for _, number := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(number)
			values = append(values, n)
		}
		numbers = append(numbers, values)
	}

	sum := 0
	for i := 0; i < len(expectations); i++ {
		expectation := expectations[i]
		values := numbers[i]

		permutations := GetPermutations(values, operators)
		if compute(values, permutations, expectation) {
			sum += expectation
		}
	}

	return strconv.Itoa(sum), nil
}
