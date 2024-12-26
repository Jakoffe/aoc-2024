package day13

import (
	"aoc/src/utils"
	"regexp"
	"strconv"
)

type Game struct {
	Ax     int
	Ay     int
	Bx     int
	By     int
	PrizeX int
	PrizeY int
}

func ReadInput(lines string, add int) []Game {
	games := []Game{}
	re := regexp.MustCompile(`(?m)X\+(\d+), Y\+(\d+)\n.*X\+(\d+), Y\+(\d+)\n.*X=(\d+), Y=(\d+)`)

	for _, match := range re.FindAllStringSubmatch(lines, -1) {
		x1, _ := strconv.Atoi(match[1])
		y1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y2, _ := strconv.Atoi(match[4])

		priceX, _ := strconv.Atoi(match[5])
		priceY, _ := strconv.Atoi(match[6])

		games = append(games, Game{
			Ax:     x1,
			Ay:     y1,
			Bx:     x2,
			By:     y2,
			PrizeX: priceX + add,
			PrizeY: priceY + add,
		})
	}

	return games
}

func Solve(game Game) (int, bool) {
	D, Dx, Dy := game.Ax*game.By-game.Bx*game.Ay, game.PrizeX*game.By-game.Bx*game.PrizeY, game.Ax*game.PrizeY-game.PrizeX*game.Ay
	if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
		return (Dx/D)*3 + (Dy / D), true
	}

	return 0, false
}

func SolveProblem1() (string, error) {
	games := ReadInput(utils.GetInput(13), 0)

	result := 0
	for _, game := range games {
		res, ok := Solve(game)

		if !ok {
			continue
		}

		result += res
	}

	return strconv.Itoa(result), nil
}

func SolveProblem2() (string, error) {
	games := ReadInput(utils.GetInput(13), 10000000000000)

	result := 0
	for _, game := range games {
		res, ok := Solve(game)

		if !ok {
			continue
		}

		result += res
	}

	return strconv.Itoa(result), nil
}
