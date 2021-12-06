package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

func main() {
	input := util.OpenFileAsString("input.txt")
	numbers := make([]int8, 0)
	for _, num := range strings.Split(strings.TrimSpace(input), ",") {
		numbers = append(numbers, int8(util.ConvertToInt(num)))
	}

	util.Start()
	fmt.Println("Part 1")
	part1(numbers)
	util.Elapsed()

	util.Start()
	fmt.Println("Part 2")
	part2(numbers)
	util.Elapsed()

	os.Exit(0)
}

func part1(input []int8) {
	fmt.Println("Result: ", iterate(input, 80))
}

func part2(input []int8) {
	fmt.Println("Result: ", iterate(input, 256))
}

func iterate(input []int8, numberOfDays int) int {
	school := make(map[int8]int)
	for _, num := range input {
		school[num]++
	}

	for i := 0; i < numberOfDays; i++ {
		born := school[0]
		school[0] = school[1]
		school[1] = school[2]
		school[2] = school[3]
		school[3] = school[4]
		school[4] = school[5]
		school[5] = school[6]
		school[6] = school[7]
		school[7] = school[8]
		school[8] = born
		school[6] += born
	}

	result := 0
	for _, v := range school {
		result += v
	}

	return result
}
