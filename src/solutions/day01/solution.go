package day01

import (
	"aoc/src/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func LoadInput(day int) ([]int, []int) {
	lines := utils.GetInput(1)

	var list1 []int
	var list2 []int

	for idx, line := range strings.Fields(lines) {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if idx%2 == 0 {
			list1 = append(list1, number)
		} else {
			list2 = append(list2, number)
		}
	}

	return list1, list2
}

func SolveProblem1() (string, error) {
	list1, list2 := LoadInput(1)
	sort.Ints(list1)
	sort.Ints(list2)

	pairs := utils.Zip(list1, list2)

	total_diff := 0.0
	for _, pair := range pairs {
		total_diff += math.Abs(float64(pair.First - pair.Second))
	}

	return strconv.Itoa(int(total_diff)), nil
}

func SolveProblem2() (string, error) {
	list1, list2 := LoadInput(1)

	counter := map[int]int{}

	simularity_score := 0
	for _, number := range list1 {
		if _, ok := counter[number]; !ok {
			counter[number] = 0
			for _, other := range list2 {
				if number == other {
					counter[number]++
				}
			}
		}

		fmt.Println("Number", number, "counter", counter[number])
		simularity_score += number * counter[number]
	}

	return strconv.Itoa(simularity_score), nil
}
