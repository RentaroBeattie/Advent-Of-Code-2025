package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var total int = 0

	//open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//part 1
	scanner := bufio.NewScanner(file)
	first := true
	var prev []byte
	//setup for part 2
	var p2file [][]byte
	//scan file line by line
	for scanner.Scan() {
		//cast to slice of bytes for mutability
		cur := []byte(scanner.Text())
		//ignores first line to fill prev
		if first {
			first = false
		} else {
			for k, v := range prev {
				//locate start / beam to
				if v == '|' || v == 'S' {
					if cur[k] == '^' {
						total++
						cur[k-1], cur[k+1] = '|', '|'
					} else {
						cur[k] = '|'
					}
				}
			}
		}
		prev = cur
		//setup for part2
		p2file = append(p2file, cur)
	}
	fmt.Println(total)

	//part 2
	//a hashmap to memoize all paths already taken
	calcedPaths := make(map[[2]int]int)
	total2 := walk(2, strings.Index(string(p2file[0]), "S"), p2file, calcedPaths)
	fmt.Println(total2)
}

func walk(row int, col int, lines [][]byte, calcedPaths map[[2]int]int) int {
	//paths that have already been taken can be returned without recomputing
	key := [2]int{row, col}
	val, found := calcedPaths[key]
	if found {
		return val
	}

	ttotal := 1
	//find splitter
	for row < len(lines)-1 && lines[row][col] != '^' {
		row++
	}
	//calculate paths from this splitter
	if row < len(lines) && lines[row][col] == '^' {
		ttotal = walk(row+2, col-1, lines, calcedPaths) + walk(row+2, col+1, lines, calcedPaths)
	}
	//memoize
	calcedPaths[key] = ttotal
	return ttotal
}
