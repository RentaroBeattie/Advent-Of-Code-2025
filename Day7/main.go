package main

import (
	"bufio"
	"fmt"
	"os"
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
	}
	fmt.Println(total)
}
