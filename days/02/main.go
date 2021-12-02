package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

func main() {
	util.Start()
	fmt.Println("Part 1")

	input := util.OpenFile("input.txt")
	part1(input)
	util.Elapsed()

	util.Start()
	fmt.Println("Part 2")
	part2(input)
	util.Elapsed()

	os.Exit(0)
}

func part1(input []string) {
	var h, d int
	for _, line := range input {
		split := strings.Split(line, " ")
		val, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			h += val
		case "up":
			d -= val
		case "down":
			d += val
		}
	}

	fmt.Println("Result: ", h*d)
}

func part2(input []string) {
	var h, depth, aim int
	for _, line := range input {
		split := strings.Split(line, " ")
		val, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			h += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}

	fmt.Println("Result: ", h*depth)
}
