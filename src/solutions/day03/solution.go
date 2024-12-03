package day03

import (
	"aoc/src/utils"
	"regexp"
	"strconv"
)

func SolveProblem1() (string, error) {
	lines := utils.GetInput(3)

	re := regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)`)

	product := 0
	for _, match := range re.FindAllStringSubmatch(lines, -1) {
		n1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		n2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		product += n1 * n2
	}

	return strconv.Itoa(product), nil
}

func SolveProblem2() (string, error) {
	lines := utils.GetInput(3)

	re := regexp.MustCompile(`(?m)mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	product := 0
	do := true
	for _, match := range re.FindAllStringSubmatch(lines, -1) {
		if match[0] == "do()" {
			do = true
			continue
		}
		if match[0] == "don't()" {
			do = false
			continue
		}

		n1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		n2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		if do {
			product += n1 * n2
		}
	}

	return strconv.Itoa(product), nil
}
