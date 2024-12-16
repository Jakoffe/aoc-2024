package day10

import (
	"aoc/src/utils"
	"fmt"
	"strconv"
	"strings"
)

var debugMode = false

type Position struct {
	x int
	y int
}

func isValidGoal(currentPosition Position, visitedGoals *[]Position, matrix *utils.Matrix[int]) bool {
	if matrix.Get(currentPosition.x, currentPosition.y) != 9 {
		return false
	}

	for _, goal := range *visitedGoals {
		if currentPosition.x == goal.x && currentPosition.y == goal.y {
			return false
		}
	}

	return true
}

func IsValidMove(currentPosition Position, newPosition Position, matrix *utils.Matrix[int]) bool {
	// Check if the new position is valid
	if newPosition.x < 0 || newPosition.y < 0 || newPosition.x >= matrix.Width() || newPosition.y >= matrix.Height() {
		return false
	}

	return matrix.Get(newPosition.x, newPosition.y) == matrix.Get(currentPosition.x, currentPosition.y)+1
}

func FindTrail(currentPosition Position, visitedGoals *[]Position, matrix *utils.Matrix[int], part2 bool) int {
	// Give debug output
	utils.DebugPrintF("Current position: %v => %d \n", debugMode, currentPosition, matrix.Get(currentPosition.x, currentPosition.y))
	if debugMode {
		matrix.PrintMapWithPosition(currentPosition.x, currentPosition.y)
	}

	// Reached the goal
	if isValidGoal(currentPosition, visitedGoals, matrix) {
		utils.DebugPrintln("GOAL REACHED", debugMode)
		if !part2 {
			*visitedGoals = append(*visitedGoals, currentPosition)
		}
		return 1
	}

	moves := []Position{}
	m1 := Position{currentPosition.x + 1, currentPosition.y}
	m2 := Position{currentPosition.x - 1, currentPosition.y}
	m3 := Position{currentPosition.x, currentPosition.y + 1}
	m4 := Position{currentPosition.x, currentPosition.y - 1}

	if IsValidMove(currentPosition, m1, matrix) {
		moves = append(moves, m1)
	}
	if IsValidMove(currentPosition, m2, matrix) {
		moves = append(moves, m2)
	}
	if IsValidMove(currentPosition, m3, matrix) {
		moves = append(moves, m3)
	}
	if IsValidMove(currentPosition, m4, matrix) {
		moves = append(moves, m4)
	}

	if len(moves) == 0 {
		utils.DebugPrintln("DEAD END", debugMode)
		return 0
	}

	if len(moves) == 1 {
		return FindTrail(moves[0], visitedGoals, matrix, part2)
	}

	utils.DebugPrintln("BRANCHING OUT", debugMode)
	trailHeads := 0
	for _, move := range moves {
		trailHeads += FindTrail(move, visitedGoals, matrix, part2)
	}
	return trailHeads
}

func SolveProblem1() (string, error) {
	lines := utils.GetInput(10)
	_ = lines

	matrix := utils.NewMatrix[int]()
	for _, line := range strings.Split(lines, "\n") {
		row := []int{}
		for _, value := range strings.Split(line, "") {
			num, _ := strconv.Atoi(value)
			row = append(row, num)
		}
		matrix.AddRow(row)
	}

	trailHeads := 0
	startingPoistions := utils.Map(matrix.FindAll(0), func(match utils.Match[int]) Position {
		return Position{match.Row, match.Column}
	})

	for _, startingPoistion := range startingPoistions {
		visitedGoals := []Position{}
		res := FindTrail(startingPoistion, &visitedGoals, matrix, false)
		fmt.Printf("Found trail heads: %v from %v\n", res, startingPoistion)
		trailHeads += res
		// break
	}

	return strconv.Itoa(trailHeads), nil
}

func SolveProblem2() (string, error) {
	lines := utils.GetInput(10)
	_ = lines

	matrix := utils.NewMatrix[int]()
	for _, line := range strings.Split(lines, "\n") {
		row := []int{}
		for _, value := range strings.Split(line, "") {
			num, _ := strconv.Atoi(value)
			row = append(row, num)
		}
		matrix.AddRow(row)
	}

	trailHeads := 0
	startingPoistions := utils.Map(matrix.FindAll(0), func(match utils.Match[int]) Position {
		return Position{match.Row, match.Column}
	})

	for _, startingPoistion := range startingPoistions {
		visitedGoals := []Position{}
		res := FindTrail(startingPoistion, &visitedGoals, matrix, true)
		fmt.Printf("Found trail heads: %v from %v\n", res, startingPoistion)
		trailHeads += res
		// break
	}

	return strconv.Itoa(trailHeads), nil
}
