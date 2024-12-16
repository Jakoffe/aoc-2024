package day11

import (
	"aoc/src/utils"
	"strconv"
	"strings"
)

var debugMode = false

type StoneArray struct {
	Stones []int
}

type Search struct {
	Stone  int
	Blinks int
}

func (s StoneArray) String() string {
	result := strings.Join(utils.Map(s.Stones, func(stone int) string {
		return strconv.Itoa(stone)
	}), ",")

	return result
}

func Blink(stone int, blinks int, cache map[Search]int) int {
	if blinks == 0 {
		return 1
	}

	var res int
	search := Search{stone, blinks}
	if value, ok := cache[search]; ok {
		return value
	} else if stone == 0 { // Rule 1
		res = Blink(1, blinks-1, cache)
	} else if len(strconv.Itoa(stone))%2 == 0 { // Rule 2
		str := strconv.Itoa(stone)
		mid := len(str) / 2

		newStone1, _ := strconv.Atoi(str[:mid])
		newStone2, _ := strconv.Atoi(str[mid:])

		res = Blink(newStone1, blinks-1, cache) + Blink(newStone2, blinks-1, cache)
	} else { // Rule 3
		res = Blink(stone*2024, blinks-1, cache)
	}

	cache[search] = res
	return res
}

func LoadStoneArray(input string) *StoneArray {
	stoneArray := StoneArray{Stones: utils.Map(strings.Split(input, " "), func(item string) int {
		num, _ := strconv.Atoi(item)
		return num
	})}

	return &stoneArray
}

func SolveProblem1() (string, error) {
	stoneArray := LoadStoneArray(utils.GetInput(11))

	cache := make(map[Search]int)
	res := utils.Sum(stoneArray.Stones, func(stone int) int {
		return Blink(stone, 25, cache)
	})

	return strconv.Itoa(res), nil
}

func SolveProblem2() (string, error) {
	stoneArray := LoadStoneArray(utils.GetInput(11))

	cache := map[Search]int{}
	res := utils.Sum(stoneArray.Stones, func(stone int) int {
		return Blink(stone, 75, cache)
	})

	return strconv.Itoa(res), nil
}
