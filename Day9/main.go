package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func FindSize(a Coord, b Coord) int {
	return (Max(a.x, b.x) - Min(a.x, b.x) + 1) * (Max(a.y, b.y) - Min(a.y, b.y) + 1)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//var total int = 0

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\r\n")
	coords := make([]Coord, 0, len(lines))
	for _, v := range lines {
		tstr := strings.Split(v, ",")
		num1, err := strconv.Atoi(tstr[0])
		num2, err2 := strconv.Atoi(tstr[1])
		if err != nil || err2 != nil {
			panic(err)
		}
		coords = append(coords, Coord{x: num1, y: num2})
	}

	largest := 0
	for _, i := range coords {
		for _, j := range coords {
			v := FindSize(i, j)
			if v > largest {
				largest = v
			}
		}
	}

	fmt.Print(largest)
}
