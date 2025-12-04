package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var total int = 0
	var total2 int = 0

	//read in files
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	//split into its lines
	lines := strings.Split(string(file), "\r\n")

	//display outcome like the example given on page
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//a slice of xy coordinates accesible rolls to remove after each iteration
	toRemove := make([][2]int, 0, 1600)

	//part 1
	total = Sweep(&lines, false, f, &toRemove)

	//part 2
	for {
		total2 += Sweep(&lines, true, f, &toRemove)

		//if no more accesible rolls exist then end loop
		if len(toRemove) == 0 {
			break
		}

		//"remove" accesible rolls
		for _, v := range toRemove {
			b := []byte(lines[v[1]])
			b[v[0]] = '.'
			lines[v[1]] = string(b)
		}
		//clear toRemove slice
		toRemove = toRemove[:0]
		//for readability in output file
		f.WriteString("\n")
	}

	fmt.Print(total)
	fmt.Print(total2)
}

func Sweep(lines *[]string, second bool, f *os.File, toRemove *[][2]int) int {
	ttotal := 0
	//iterate through each byte in each line
	for y, line := range *lines {
		for x, v := range line {
			if v == '@' && CountNeighbours(x, y, lines) < 4 {
				_, err := f.WriteString("X")
				if err != nil {
					panic(err)
				}
				if second {
					//add coordinate of accesible roll
					vector := [2]int{x, y}
					*toRemove = append((*toRemove), vector)
				}
				//increment total found
				ttotal++
			} else {
				_, err := f.WriteString(string(v))
				if err != nil {
					panic(err)
				}
			}
		}
		f.WriteString("\n")
	}

	return ttotal
}

func CountNeighbours(x int, y int, lines *[]string) int {
	//check all adjacent tiles
	count := Check(x-1, y, lines)
	count += Check(x+1, y, lines)
	count += Check(x, y-1, lines)
	count += Check(x, y+1, lines)
	count += Check(x-1, y+1, lines)
	count += Check(x+1, y+1, lines)
	count += Check(x-1, y-1, lines)
	count += Check(x+1, y-1, lines)
	return count
}

func Check(x int, y int, lines *[]string) int {
	//validate if it is '@' and within bounds
	if x >= 0 && x < len((*lines)[0]) && y >= 0 && y < len(*lines) && (*lines)[y][x] == '@' {
		return 1
	}
	return 0
}
