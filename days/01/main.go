package main

import (
	"fmt"
	"os"

	util "github.com/madimaa/aoc2021/util"
)

func main() {
	input := util.OpenFileAsIntArray("input.txt")

	util.Start()
	fmt.Println("Part 1")

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
