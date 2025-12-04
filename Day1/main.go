package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WrapAround(n int, dial int) int {
	v := (dial + n) % 100
	if v < 0 {
		v += 100
	}
	return v
}

func ClickThrough(n int, dial int, passes int) int {
	v := dial + n

	x := v / 100

	if n > 0 {
		passes += x
	} else {
		//land on 0 edge case
		if v%100 == 0 {
			passes++
		}
		//anull above if started on 0
		if dial == 0 {
			passes--
		}
		for v < 0 {
			v += 100
			passes++
		}
	}

	return passes
}

func main() {
	//initialise values
	var password int = 0
	var password2 int = 0
	var dial int = 50

	//open file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sign, val := 1, 0

	//start scanning the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//determine direction
		if line[0] == 82 {
			sign = 1
		} else {
			sign = -1
		}

		val, err = strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		//part 2
		password2 = ClickThrough(sign*val, dial, password2)

		//part 1
		dial = WrapAround(sign*val, dial)

		if dial == 0 {
			password += 1
		}

	}

	fmt.Println(password)
	fmt.Println(password2)
}
