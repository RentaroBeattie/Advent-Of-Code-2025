package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int = 0

	//read in file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	//split into rows of strings
	raw := strings.Split(string(file), "\r\n")
	//parse integers
	var lines [][]string
	for _, v := range raw {
		lines = append(lines, strings.Fields(v))
	}

	//iterate  through each column and calculate sums
	for i := 0; i < len(lines[0]); i++ {
		vals := make([]int, 0, len(lines)-1)
		for j := 0; j < len(lines)-1; j++ {
			temp, err := strconv.Atoi(lines[j][i])
			if err != nil {
				panic(err)
			}
			vals = append(vals, temp)
		}
		switch lines[len(lines)-1][i] {
		case "+": //addition case
			total += calculate(vals, func(x int, y int) int { return x + y })
		case "*": //multiplication case
			total += calculate(vals, func(x int, y int) int { return x * y })
		}
	}
	fmt.Println(total)

	//part 2
	var total2 int
	var nums []int
	numctr := 0
	//iterate through each byte column
	for i := 0; i < len(raw[0]); i++ {
		rawstr := ""
		null := 0
		//merge if not whitespace
		for j := 0; j < len(raw)-1; j++ {
			if raw[j][i] == ' ' {
				null++
			} else {
				rawstr = rawstr + string(raw[j][i])
			}
		}
		//calculate values
		if null == len(raw)-1 {
			switch lines[len(lines)-1][numctr] {
			case "+":
				total2 += calculate(nums, func(x int, y int) int { return x + y })
			case "*":
				total2 += calculate(nums, func(x int, y int) int { return x * y })
			}
			nums = nums[:0]
			numctr++
		} else {
			temp, err := strconv.Atoi(rawstr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, temp)
		}
	}
	//handle final column edge case
	switch lines[len(lines)-1][numctr] {
	case "+":
		total2 += calculate(nums, func(x int, y int) int { return x + y })
	case "*":
		total2 += calculate(nums, func(x int, y int) int { return x * y })
	}

	fmt.Print(total2)
}

// calculate the sums of all values in a columns with the passed in function
func calculate(vals []int, operation func(x int, y int) int) int {
	sum := vals[0]
	for i := 1; i < len(vals); i++ {
		sum = operation(sum, vals[i])
	}
	return sum
}
