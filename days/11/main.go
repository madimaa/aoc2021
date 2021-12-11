package main

import (
	"fmt"
	"os"

	util "github.com/madimaa/aoc2021/util"
)

type pair struct {
	x, y int
}

func main() {
	input := util.OpenFile("input.txt")

	util.Start()
	fmt.Println("Part 1")
	part1(input)
	util.Elapsed()

	util.Start()
	fmt.Println("Part 2")
	part2(input)
	util.Elapsed()

	os.Exit(0)
}

func part1(input []string) {
	xMax, yMax := len(input[0]), len(input)
	grid := Create(xMax, yMax)
	for y, line := range input {
		for x, r := range line {
			grid.Put(x, y, util.ConvertToInt(string(r)))
		}
	}

	flashes := 0
	for i := 0; i < 100; i++ {
		grid.increase()
		flashes += grid.getFlashes()
		grid.zeroFlashes()
	}

	fmt.Println("Result: ", flashes)
}

func part2(input []string) {
	xMax, yMax := len(input[0]), len(input)
	grid := Create(xMax, yMax)
	for y, line := range input {
		for x, r := range line {
			grid.Put(x, y, util.ConvertToInt(string(r)))
		}
	}

	counter := 0
	for {
		counter++
		grid.increase()
		if grid.getFlashes() == 100 {
			break
		}
		grid.zeroFlashes()

	}

	fmt.Println("Result: ", counter)
}
