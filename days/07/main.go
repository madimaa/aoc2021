package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

func main() {
	input := util.OpenFileAsString("input.txt")
	numbers := make([]int, 0)
	for _, num := range strings.Split(strings.TrimSpace(input), ",") {
		numbers = append(numbers, util.ConvertToInt(num))
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

func part1(input []int) {
	min := 0
	max := input[0]
	for _, val := range input {
		if min > val {
			min = val
		} else if max < val {
			max = val
		}
	}

	minfuel := 0
	for i := min; i <= max; i++ {

		fuel := 0
		for _, val := range input {
			fuel += int(math.Abs(float64(i - val)))
		}

		if minfuel == 0 || minfuel > fuel {
			minfuel = fuel
		}
	}

	fmt.Println("Result: ", minfuel)
}

func part2(input []int) {
	min := 0
	max := input[0]
	for _, val := range input {
		if min > val {
			min = val
		} else if max < val {
			max = val
		}
	}

	fuelcost := make(map[int]int)
	fuelcost[1] = 1
	for i := 2; i <= max-min; i++ {
		fuelcost[i] = fuelcost[i-1] + i
	}

	minfuel := 0
	for i := min; i <= max; i++ {

		fuel := 0
		for _, val := range input {
			fuel += fuelcost[int(math.Abs(float64(i-val)))]
		}

		if minfuel == 0 || minfuel > fuel {
			minfuel = fuel
		}
	}

	fmt.Println("Result: ", minfuel)
}
