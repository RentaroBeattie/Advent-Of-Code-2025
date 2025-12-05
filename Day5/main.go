package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Idrange struct {
	upper int
	lower int
}

func main() {
	var total int = 0

	//read in input.txt
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	//split the file down, removing carriage returns and newline
	lines := strings.Split(string(file), "\r\n")

	//determine the index of the dividing line
	seperator := 0
	for k, v := range lines {
		if v == "" {
			seperator = k
			break
		}
	}

	//parse ranges into a slice of structs for easier access
	idRanges := make([]Idrange, 0, 200)
	for i := 0; i < seperator; i++ {
		//split into lower and upper bounds
		rng := strings.Split(lines[i], "-")
		rngl, err1 := strconv.Atoi(rng[0])
		rngu, err2 := strconv.Atoi(rng[1])
		if err1 != nil || err2 != nil {
			panic(err1)
		}
		idRanges = append(idRanges, Idrange{upper: rngu, lower: rngl})
	}

	//iterate through each ingredient
	for i := seperator + 1; i < len(lines); i++ {
		id, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}
		//iterate through each fresh ingredient id ranges
		for j := 0; j < seperator; j++ {
			//increment total and break loop as ingredient id is fresh if it is in any range
			if id >= idRanges[j].lower && id <= idRanges[j].upper {
				total++
				break
			}
		}
	}

	//Part 2
	var total2 int = 0

	//cut down colliding ranges to avoid duplicate ids
	finalIds := make([]Idrange, 0, 200)
	for i := 0; i < len(idRanges); i++ {
		//fmt.Printf("%v %v", i, len(finalIds))
		cur := idRanges[i]
		//compare to all existing ranges to prevent collisions
		j := 0
		newFinal := make([]Idrange, 0, len(finalIds)+1)
	outer:
		for j < len(finalIds) {
			v := finalIds[j]
			switch {
			case cur.lower < v.lower && cur.upper <= v.upper: //right collision - trim
				cur.upper = v.lower - 1
				newFinal = append(newFinal, v)
				j++
			case cur.lower >= v.lower && cur.upper > v.upper: //left collision - trim
				cur.lower = v.upper + 1
				newFinal = append(newFinal, v)
				j++
			case cur.lower >= v.lower && cur.upper <= v.upper: //inner collision - invalidate current range
				cur.lower, cur.upper = -1, -1
				newFinal = append(newFinal, v)
				break outer
			case cur.lower < v.lower && cur.upper > v.upper: //super size - remove original
				j++
			default:
				newFinal = append(newFinal, v)
				j++
			}
			//fmt.Printf("%v %v - %v %v\n", cur.lower, cur.upper, v.lower, v.upper)
		}

		finalIds = newFinal
		//if isnt invalid range then append to finalIds
		if cur.lower != -1 && cur.upper != -1 {
			finalIds = append(finalIds, cur)
		}
		//fmt.Print(" Done\n")
	}

	//increment total2 by each non-colliding ranges (inclusive)
	for _, v := range finalIds {
		total2 += v.upper - v.lower + 1
	}

	fmt.Println(total)
	fmt.Println(total2)
}
