package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int = 0
	var total2 int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		max_val, max_ind := 0, 0

		for k, v := range line { //search for the largest digit in the bank excluding the final digit
			val := int(v - '0')
			if val > max_val && k != len(line)-1 {
				max_val = val
				max_ind = k
			}
		}

		nmax_ind := max_ind + 1 //search for the largest digit after the first largest
		nmax_val := int(line[nmax_ind] - '0')
		for i := max_ind + 2; i < len(line); i++ {
			if int(line[i]-'0') > nmax_val {
				nmax_val = int(line[i] - '0')
				nmax_ind = i
			}
		}

		total += max_val*10 + nmax_val

		digits := make([]int, 0, 12)
		Part2(line, &digits, 0)

		mult := 1
		for i := 11; i >= 0; i-- {
			total2 += digits[i] * mult
			mult *= 10
		}
	}

	fmt.Println(total)
	fmt.Println(total2)
}

func Part2(line string, digits *[]int, index int) {
	dmax_val := int(line[index] - '0')
	dmax_ind := index
	for i := index; i < len(line)-11+len(*digits); i++ {
		if int(line[i]-'0') > dmax_val {
			dmax_val = int(line[i] - '0')
			dmax_ind = i
		}
	}
	*digits = append(*digits, dmax_val)
	if len(*digits) < cap(*digits) {
		Part2(line, digits, dmax_ind+1)
	}
}
