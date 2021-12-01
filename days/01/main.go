package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2021/util"
)

func main() {
	util.Start()
	fmt.Println("Part 1")

	file, err := os.Open("input.txt")
	util.PanicOnError(err)

	scanner := bufio.NewScanner(file)
	input := make([]int, 0)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		util.LogOnError(err)
		input = append(input, number)
	}

	util.LogOnError(scanner.Err())
	util.LogOnError(file.Close())

	inc := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			inc++
		}
	}

	fmt.Println("Result: ", inc)
	util.Elapsed()

	util.Start()
	fmt.Println("Part 2")

	inc = 0
	for i := 1; i < len(input)-2; i++ {
		a := input[i-1] + input[i] + input[i+1]
		b := input[i] + input[i+1] + input[i+2]
		if a < b {
			inc++
		}
	}

	fmt.Println("Result: ", inc)
	util.Elapsed()
	os.Exit(0)
}
