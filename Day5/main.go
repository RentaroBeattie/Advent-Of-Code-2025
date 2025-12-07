package main

import (
	"fmt"
	"os"
	"sort"
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

	//Sort ranges by lower bound
	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i].lower < idRanges[j].lower
	})

	//Merge overlapping/adjacent ranges
	merged := []Idrange{idRanges[0]}
	for i := 1; i < len(idRanges); i++ {
		cur := idRanges[i]
		last := &merged[len(merged)-1]

		if cur.lower <= last.upper+1 { // Overlapping or adjacent
			last.upper = max(last.upper, cur.upper)
		} else {
			merged = append(merged, cur)
		}
	}

	//Count total IDs in merged ranges
	for _, v := range merged {
		total2 += v.upper - v.lower + 1
	}

	fmt.Println(total)
	fmt.Println(total2)
}
