package day02

import (
	"aoc/src/utils"
	"math"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}

func IsSafeWithoutDamper(levels []int) bool {
	diffs := []int{}

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		diffs = append(diffs, diff)
	}

	increasing := true
	decreasing := true

	for _, diff := range diffs {
		abs := math.Abs(float64(diff))

		if abs < 1 || abs > 3 {
			return false
		}

		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}

func (r Report) IsSafe(withDamper bool) bool {
	if withDamper {
		for i := 0; i < len(r.Levels); i++ {
			slice := append(append([]int{}, r.Levels[:i]...), r.Levels[i+1:]...)

			if IsSafeWithoutDamper(slice) {
				return true
			}
		}
		return false
	}

	return IsSafeWithoutDamper(r.Levels)
}

func SolveProblem1() (string, error) {
	lines := utils.GetInput(2)
	reportStrings := strings.Split(lines, "\n")

	safeReports := 0

	for _, reportString := range reportStrings {
		if len(reportString) == 0 {
			continue
		}

		var levels []int
		for _, level := range strings.Split(reportString, " ") {
			number, err := strconv.Atoi(level)

			if err != nil {
				panic(err)
			}

			levels = append(levels, number)
		}

		report := Report{Levels: levels}

		if report.IsSafe(false) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports), nil
}

func SolveProblem2() (string, error) {
	lines := utils.GetInput(2)
	reportStrings := strings.Split(lines, "\n")

	safeReports := 0

	for _, reportString := range reportStrings {
		if len(reportString) == 0 {
			continue
		}

		var levels []int
		for _, level := range strings.Split(reportString, " ") {
			number, err := strconv.Atoi(level)

			if err != nil {
				panic(err)
			}

			levels = append(levels, number)
		}

		report := Report{Levels: levels}

		if report.IsSafe(true) {
			safeReports++
		}
	}

	return strconv.Itoa(safeReports), nil
}
