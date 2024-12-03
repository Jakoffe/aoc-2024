package main

//go:generate go run ./gen

import (
	"aoc/src/solutions/day01"
	"aoc/src/solutions/day02"
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

var DaySelected int
var PartSelected int

type aocResult struct {
	Result      string
	Error       error
	TimeElapsed time.Duration
}

func init() {
	pflag.IntVarP(&DaySelected, "day", "d", 1, "run specific day")
	pflag.IntVarP(&PartSelected, "part", "p", 1, "run specific part")
}

func main() {
	pflag.Parse()

	result := runAocPart(DaySelected, PartSelected)

	fmt.Println("Day", DaySelected, "part", PartSelected)
	fmt.Println("Result:", result.Result)
	fmt.Println("Time elapsed:", result.TimeElapsed)
}

func runAocPart(day int, part int) aocResult {
	fmt.Println("Running day", day, "part", part)
	start := time.Now()

	var res string
	var err error

	switch day {
	case 1:
		if part == 1 {
			res, err = day01.SolveProblem1()
		} else if part == 2 {
			res, err = day01.SolveProblem2()
		}
	case 2:
		if part == 1 {
			res, err = day02.SolveProblem1()
		} else if part == 2 {
			res, err = day02.SolveProblem2()
		}
	default:
		panic(fmt.Sprintf("Unknown day %d", day))
	}

	return aocResult{Result: res, Error: err, TimeElapsed: time.Since(start)}
}