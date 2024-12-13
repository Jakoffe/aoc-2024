package day08

import (
	"aoc/src/utils"
	"strconv"
	"strings"
)

type AntennaPair struct {
	First  utils.Match[string]
	Second utils.Match[string]
}

func LoadInput() utils.Matrix[string] {
	lines := strings.Split(utils.GetInput(8), "\n")
	matrix := utils.NewMatrix[string]()
	for _, line := range lines {
		matrix.AddRow(strings.Split(line, ""))
	}

	return *matrix
}

func computeAntennaPairs(matrix utils.Matrix[string], uniqueAntennas utils.Set[string]) []AntennaPair {
	pairs := []AntennaPair{}
	for _, uniqueAntenna := range uniqueAntennas.Members() {
		antennas := matrix.FindAll(uniqueAntenna)

		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				pairs = append(pairs, AntennaPair{antennas[i], antennas[j]})
			}
		}

	}

	return pairs
}

func (pair AntennaPair) computeAntiNodesV1(matrix utils.Matrix[string]) {
	xdiff := pair.First.Row - pair.Second.Row
	ydiff := pair.First.Column - pair.Second.Column

	if xdiff == 0 && ydiff == 0 {
		panic("xdiff and ydiff are both 0")
	}

	matrix.Set(pair.First.Row+xdiff, pair.First.Column+ydiff, "#")
	matrix.Set(pair.Second.Row-xdiff, pair.Second.Column-ydiff, "#")
}

func (pair AntennaPair) computeAntiNodesV2(matrix utils.Matrix[string]) {
	xdiff := pair.First.Row - pair.Second.Row
	ydiff := pair.First.Column - pair.Second.Column

	x1 := pair.First.Row + xdiff
	y1 := pair.First.Column + ydiff
	for matrix.Set(x1, y1, "#") {
		x1 += xdiff
		y1 += ydiff
	}

	x2 := pair.Second.Row - xdiff
	y2 := pair.Second.Column - ydiff
	for matrix.Set(x2, y2, "#") {
		x2 -= xdiff
		y2 -= ydiff
	}
}

func SolveProblem1() (string, error) {
	matrix := LoadInput()

	uniqueAntennas := utils.NewSet(matrix.Elements()...).Remove(".")
	println("Found unique antennas:", uniqueAntennas.String())

	pairs := computeAntennaPairs(matrix, uniqueAntennas)

	for _, pair := range pairs {
		pair.computeAntiNodesV1(matrix)
	}

	result := len(matrix.FindAll("#"))

	return strconv.Itoa(result), nil
}

func SolveProblem2() (string, error) {
	matrix := LoadInput()

	uniqueAntennas := utils.NewSet(matrix.Elements()...).Remove(".")
	println("Found unique antennas:", uniqueAntennas.String())

	pairs := computeAntennaPairs(matrix, uniqueAntennas)

	for _, pair := range pairs {
		pair.computeAntiNodesV2(matrix)
	}

	result := matrix.Height()*matrix.Width() - len(matrix.FindAll("."))

	return strconv.Itoa(result), nil
}
