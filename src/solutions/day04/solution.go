package day04

import (
	"aoc/src/utils"
	"strconv"
	"strings"
)

func checkRow(matrix utils.Matrix[string], i int, j int, searchStrings []string, stringLength int) int {
	if j+stringLength > matrix.Width() {
		return 0
	}

	rowSlice := matrix.GetRow(i)[j : j+stringLength]
	if matches := checkSearchStrings(strings.Join(rowSlice, ""), searchStrings); matches > 0 {
		return matches
	}

	return 0
}

func checkCol(matrix utils.Matrix[string], i int, j int, searchStrings []string, stringLength int) int {
	if i+stringLength > matrix.Height() {
		return 0
	}

	columnSlice := matrix.GetColumn(j)[i : i+stringLength]
	if matches := checkSearchStrings(strings.Join(columnSlice, ""), searchStrings); matches > 0 {
		return matches
	}

	return 0
}

func checkDiagonal(matrix utils.Matrix[string], i int, j int, searchStrings []string, stringLength int) int {
	diagonal := matrix.GetDiagonal(i, j, false)
	if len(diagonal) < stringLength {
		return 0
	}

	diagonal = diagonal[:stringLength]
	if matches := checkSearchStrings(strings.Join(diagonal, ""), searchStrings); matches > 0 {
		return matches
	}

	return 0
}

func checkAntiDiagonal(matrix utils.Matrix[string], i int, j int, searchStrings []string, stringLength int) int {
	antiDiagonal := matrix.GetAntiDiagonal(i, j, false)
	if len(antiDiagonal) < stringLength {
		return 0
	}

	antiDiagonal = antiDiagonal[:stringLength]
	if matches := checkSearchStrings(strings.Join(antiDiagonal, ""), searchStrings); matches > 0 {
		return matches
	}

	return 0
}

func checkCross(matrix utils.Matrix[string], i int, j int, searchStrings []string, stringLength int) int {
	matches := 0
	if i < 1 || i > matrix.Height()-2 || j < 1 || j > matrix.Width()-2 {
		return matches
	}

	smallMatrix := utils.Matrix[string]{}

	smallMatrix.AddRow(matrix.GetRow(i - 1)[j-1 : j+2])
	smallMatrix.AddRow(matrix.GetRow(i)[j-1 : j+2])
	smallMatrix.AddRow(matrix.GetRow(i + 1)[j-1 : j+2])

	// Check diagonal and anti-diagonal
	if checkDiagonal(smallMatrix, 0, 0, searchStrings, stringLength) > 0 && checkAntiDiagonal(smallMatrix, 0, 2, searchStrings, stringLength) > 0 {
		matches++
	}

	return matches
}

func checkSearchStrings(slice string, searchStrings []string) int {
	matches := 0
	for _, searchString := range searchStrings {
		if slice == searchString {
			matches++
		}
	}

	return matches
}

func SolveProblem1() (string, error) {
	lines := utils.GetInput(4)
	matrix := utils.Matrix[string]{}
	searchStrings := []string{"XMAS", "SAMX"}
	stringLength := len(searchStrings[0])

	for _, line := range strings.Fields(lines) {
		matrix.AddRow(strings.Split(line, ""))
	}

	counter := 0
	for i := 0; i < matrix.Height(); i++ {
		for j := 0; j < matrix.Width(); j++ {
			if matrix.Get(i, j) == string(searchStrings[0][1]) || matrix.Get(i, j) == string(searchStrings[1][0]) {
				// Check row
				counter += checkRow(matrix, i, j, searchStrings, stringLength)

				// Check column
				counter += checkCol(matrix, i, j, searchStrings, stringLength)

				// Check diagonal
				counter += checkDiagonal(matrix, i, j, searchStrings, stringLength)

				// Check anti-diagonal
				counter += checkAntiDiagonal(matrix, i, j, searchStrings, stringLength)
			}
		}
	}

	return strconv.Itoa(counter), nil
}

func SolveProblem2() (string, error) {
	lines := utils.GetInput(4)
	matrix := utils.Matrix[string]{}
	searchStrings := []string{"MAS", "SAM"}
	stringLength := len(searchStrings[0])

	for _, line := range strings.Fields(lines) {
		matrix.AddRow(strings.Split(line, ""))
	}

	counter := 0
	for i := 0; i < matrix.Height(); i++ {
		for j := 0; j < matrix.Width(); j++ {
			if matrix.Get(i, j) == "A" {
				// Check row and column
				counter += checkCross(matrix, i, j, searchStrings, stringLength)
			}
		}
	}

	return strconv.Itoa(counter), nil
}
