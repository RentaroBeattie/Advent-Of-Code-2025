package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int = 0
	var total2 int = 0

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	initRanges := strings.Split(string(file), ",")

	for _, v := range initRanges {
		tempSplit := strings.Split(v, "-")

		first, err := strconv.Atoi(tempSplit[0])
		if err != nil {
			panic(err)
		}
		last, err := strconv.Atoi(tempSplit[1])
		if err != nil {
			panic(err)
		}

		for i := first; i <= last; i++ {
			str := strconv.Itoa(i) //convert to string
			length := len(str)
			//part 2
			for j := 1; j <= length/2; j++ { //increment through possible lengths of sub sequence
				if length%j != 0 { //if number cannot be equally partitioned with j then skip
					continue
				}
				interval := length / j
				count := 1
				for l := 1; l < interval; l++ { //compare all other subsequences to the first
					if str[:j] != str[l*j:l*j+j] { //compare first subsequence to current subsequence
						break //break if any incorrect
					}
					count++
				}
				if count == interval { //if all are equal then add to total2
					total2 += i
					break
				}
			}
			//part 1
			if length%2 == 1 { //odd lengthed numbers cannot be mirrored
				continue
			}
			if str[:length/2] == str[length/2:] { //add number to total if both halves are equals
				total += i
			}
		}
	}

	fmt.Printf("%v %v", total, total2)
}
