package day14

import (
	"aoc/src/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	X      int
	Y      int
	XDelta int
	YDelta int
}

func (r Robot) Print(height int, width int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if r.X == x && r.Y == y {
				print("1")
			} else {
				print(".")
			}
		}
		print("\n")
	}
	println("========================")
}

func GetQuadrant(robot *Robot, height int, width int) (int, bool) {
	// qx := robot.X / width / 2
	// qy := robot.Y / height / 2
	if robot.X == width/2 || robot.Y == height/2 {
		return 0, false
	}

	qx := int(math.Floor(float64(robot.X) / math.Floor(float64(width/2))))
	qy := int(math.Floor(float64(robot.Y) / math.Floor(float64(height/2))))

	// println(qx, qy)

	quad := qx
	if qx < 2 {
		quad++
	}

	if qy > 0 {
		quad += 2
	}

	return quad, true
}

func QuadrantCount(robots []*Robot, height int, width int) []int {
	QuadrantCount := []int{0, 0, 0, 0}

	for _, robot := range robots {
		quadrant, OK := GetQuadrant(robot, height, width)

		if !OK {
			continue
		}

		QuadrantCount[quadrant-1]++
	}

	return QuadrantCount
}

func PrintRobots(robots []*Robot, height int, width int) {
	counter := make([][]uint8, height)
	for i := range counter {
		counter[i] = make([]uint8, width)
	}

	for _, robot := range robots {
		counter[robot.Y][robot.X] += 1
	}

	for _, row := range counter {
		for _, count := range row {
			if count > 0 {
				print(count)
			} else {
				print(".")
			}
		}
		println()
	}
	println("========================")

}

func (r *Robot) Move(height int, width int, seconds int) {
	newX := (r.X + r.XDelta*seconds) % width
	newY := (r.Y + r.YDelta*seconds) % height

	if newX < 0 {
		newX = width + newX
	}

	if newY < 0 {
		newY = height + newY
	}

	r.X, r.Y = newX, newY
}

func LoadInput(lines string) []*Robot {
	robots := make([]*Robot, strings.Count(lines, "p="))
	re := regexp.MustCompile(`(?m)^p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for i, match := range re.FindAllStringSubmatch(lines, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		xDelta, _ := strconv.Atoi(match[3])
		yDelta, _ := strconv.Atoi(match[4])

		robots[i] = &Robot{
			X:      x,
			Y:      y,
			XDelta: xDelta,
			YDelta: yDelta,
		}
	}

	return robots
}

func SolveProblem1() (string, error) {
	robots := LoadInput(utils.GetInput(14))

	seconds := 100
	height, width := 103, 101

	for _, robot := range robots {
		robot.Move(height, width, seconds)
	}

	Quadrants := QuadrantCount(robots, height, width)

	for i, v := range Quadrants {
		fmt.Printf("Quadrant %v: %v\n", i, v)
	}

	safetyFactor := utils.Prod(Quadrants, func(v int) int {
		return v
	})

	return strconv.Itoa(safetyFactor), nil
}

func SolveProblem2() (string, error) {
	robots := LoadInput(utils.GetInput(14))
	height, width := 103, 101

	var tree int
outer:
	for i := 1; i < 10000; i++ {
		robotMap := map[utils.Coordinate]bool{}
		for _, robot := range robots {
			robot.Move(height, width, 1)
			robotMap[utils.Coordinate{X: robot.X, Y: robot.Y}] = true
		}

		for row := 0; row < height; row++ {
			consecutiveRobots := 0
			for col := 0; col < width; col++ {
				if robotMap[utils.Coordinate{X: col, Y: row}] {
					consecutiveRobots++
				} else {
					consecutiveRobots = 0
				}

				if consecutiveRobots > 10 {
					tree = i
					break outer
				}
			}

		}
	}

	return strconv.Itoa(tree), nil
}
