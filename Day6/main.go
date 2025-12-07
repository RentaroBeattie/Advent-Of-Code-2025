package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int = 0

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	raw := strings.Split(string(file), "\r\n")
	var lines [][]string
	for _, v := range raw {
		lines = append(lines, strings.Fields(v))
	}

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
		case "+":
			total += calculate(vals, func(x int, y int) int { return x + y })
		case "*":
			total += calculate(vals, func(x int, y int) int { return x * y })
		}
	}
	fmt.Println(total)

	//part 2
	var total2 int
	cephvals := make([][]int, 0, len(lines[0]))
	var nums []int
	for i := 0; i < len(raw[0]); i++ {
		fmt.Printf("%v. ", len(cephvals))
		rawstr := ""
		null := 0
		for j := 0; j < len(raw)-1; j++ {
			if raw[j][i] == ' ' {
				null++
			} else {
				rawstr = rawstr + string(raw[j][i])
			}
		}
		if null == len(raw)-1 {
			cephvals = append(cephvals, nums)
			fmt.Println(nums)
			nums = nums[:0]
		} else {
			temp, err := strconv.Atoi(rawstr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, temp)
			fmt.Printf("%v, ", temp)
		}
	}
	cephvals = append(cephvals, nums)

	for k, v := range cephvals {
		/*
			for _, v2 := range v {
				fmt.Printf("%v, ", v2)
			}
			fmt.Println(lines[len(lines)-1][k])
		*/
		switch lines[len(lines)-1][k] {
		case "+":
			total2 += calculate(v, func(x int, y int) int { return x + y })
		case "*":
			total2 += calculate(v, func(x int, y int) int { return x * y })
		}
	}
	fmt.Println(total2)
}

func calculate(vals []int, operation func(x int, y int) int) int {
	sum := vals[0]
	for i := 1; i < len(vals); i++ {
		sum = operation(sum, vals[i])
	}
	return sum
}
