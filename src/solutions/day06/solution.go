package day06

import (
	"aoc/src/utils"
	"strconv"
	"strings"
)

type Move struct {
	row         int
	column      int
	orientation string
}

type Guard struct {
	row         int
	column      int
	orientation string
}

type Lab struct {
	graph [][]string
	guard Guard
	moves []Move
}

func (l Lab) print() {
	for row, line := range l.graph {
		output := ""
		for col, char := range line {
			if row == l.guard.row && col == l.guard.column {
				output += l.guard.orientation
			} else {
				output += char
			}
		}

		println(output)
	}
}

func (l Lab) hasLoop() bool {

	for _, move := range l.moves {
		if move.row == l.guard.row && move.column == l.guard.column && move.orientation == l.guard.orientation {
			return true
		}
	}

	return false
}

func (l *Lab) updateOrientation() {
	switch l.guard.orientation {
	case "^":
		l.guard.orientation = ">"
	case ">":
		l.guard.orientation = "v"
	case "v":
		l.guard.orientation = "<"
	case "<":
		l.guard.orientation = "^"
	}
}

func (l Lab) isGuardInLab(row int, column int) bool {
	width := len(l.graph[0])
	height := len(l.graph)

	if row < 0 || row >= height {
		return false
	}

	if column < 0 || column >= width {
		return false
	}

	return true
}

func (l Lab) NextField() (int, int) {
	switch l.guard.orientation {
	case "^":
		return l.guard.row - 1, l.guard.column
	case ">":
		return l.guard.row, l.guard.column + 1
	case "v":
		return l.guard.row + 1, l.guard.column
	case "<":
		return l.guard.row, l.guard.column - 1
	}

	panic("Error")
}

func (l *Lab) performMove() bool {
	newRow, newColumn := l.NextField()
	l.moves = append(l.moves, Move{row: l.guard.row, column: l.guard.column, orientation: l.guard.orientation})

	if !l.isGuardInLab(newRow, newColumn) {
		l.guard.row = newRow
		l.guard.column = newColumn

		return true
	}

	newField := l.graph[newRow][newColumn]
	if newField == "." || newField == "X" {
		l.guard.row = newRow
		l.guard.column = newColumn
		l.graph[newRow][newColumn] = "X"
		return false
	}

	if newField == "#" {
		l.updateOrientation()
		return false
	}

	return true
}

func (l *Lab) solvable() bool {
	for l.isGuardInLab(l.guard.row, l.guard.column) {
		l.performMove()
		if l.hasLoop() {
			return false
		}
	}

	return true
}

func (l Lab) makeCopy() Lab {
	copiedGraph := make([][]string, len(l.graph))
	for i := range l.graph {
		copiedGraph[i] = make([]string, len(l.graph[i]))
		copy(copiedGraph[i], l.graph[i])
	}

	// Deep copy the guard (simple struct copy)
	copiedGuard := Guard{
		row:         l.guard.row,
		column:      l.guard.column,
		orientation: l.guard.orientation,
	}

	// Deep copy the moves slice
	copiedMoves := make([]Move, len(l.moves))
	copy(copiedMoves, l.moves)

	// Return the new Lab struct
	return Lab{
		graph: copiedGraph,
		guard: copiedGuard,
		moves: copiedMoves,
	}
}

func SolveProblem1() (string, error) {
	lines := utils.GetInput(6)
	Lab := Lab{graph: [][]string{}, guard: Guard{}}

	for x, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}

		row := []string{}
		for y, char := range strings.Split(line, "") {
			if char == "^" {
				Lab.guard = Guard{row: x, column: y, orientation: "^"}
				row = append(row, "X")
			} else {
				row = append(row, char)
			}
		}
		Lab.graph = append(Lab.graph, row)
	}

	for Lab.isGuardInLab(Lab.guard.row, Lab.guard.column) {
		Lab.performMove()
	}

	moves := 0
	for i := 0; i < len(Lab.graph); i++ {
		for j := 0; j < len(Lab.graph[i]); j++ {
			if Lab.graph[i][j] == "X" {
				moves += 1
			}
		}
	}

	return strconv.Itoa(moves), nil
}

func SolveProblem2() (string, error) {
	lines := utils.GetInput(6)
	lab := Lab{graph: [][]string{}, guard: Guard{}}

	for x, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}

		row := []string{}
		for y, char := range strings.Split(line, "") {
			if char == "^" {
				lab.guard = Guard{row: x, column: y, orientation: "^"}
				row = append(row, "X")
			} else {
				row = append(row, char)
			}
		}
		lab.graph = append(lab.graph, row)
	}

	initialLab := lab.makeCopy()
	initialLab.solvable()

	possibilities := 0
	for i := 0; i < len(lab.graph); i++ {
		for j := 0; j < len(lab.graph[i]); j++ {
			if initialLab.graph[i][j] != "X" {
				continue
			}
			testLab := lab.makeCopy()
			testLab.graph[i][j] = "#"

			if !testLab.solvable() {
				possibilities++
			}
		}
	}

	return strconv.Itoa(possibilities), nil
}
