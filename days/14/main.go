package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

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
	polymer := input[0]
	pairs := make(map[string]string)
	for i := 2; i < len(input); i++ {
		line := input[i]
		split := strings.Split(line, " -> ")
		pairs[split[0]] = split[1]
	}

	result := doSteps(10, polymer, pairs)
	counter := make(map[rune]int)
	for _, r := range result {
		counter[r]++
	}

	min := counter[rune(result[0])]
	max := 0
	for _, val := range counter {
		if max < val {
			max = val
		}

		if min > val {
			min = val
		}
	}

	fmt.Println("Result: ", max-min)
}

func part2(input []string) {
	polymer := input[0]
	pairs := make(map[string]string)
	for i := 2; i < len(input); i++ {
		line := input[i]
		split := strings.Split(line, " -> ")
		pairs[split[0]] = split[1]
	}

	counter := doStepsOptimized(40, polymer, pairs)

	min := 9999999999999999
	max := 0
	for _, val := range counter {
		if max < val {
			max = val
		}

		if min > val {
			min = val
		}
	}

	fmt.Println("Result: ", max-min)
}

func doSteps(count int, polymer string, pairs map[string]string) string {
	for i := 0; i < count; i++ {
		newPolymer := string(polymer[0])
		for i := 1; i < len(polymer); i++ {
			prev := string(polymer[i-1])
			act := string(polymer[i])
			group := prev + act
			if val, ok := pairs[group]; ok {
				newPolymer += val + act
			} else {
				newPolymer += act
			}
		}
		polymer = newPolymer
	}

	return polymer
}

func doStepsOptimized(limit int, polymer string, pairs map[string]string) map[rune]int {
	bucket := make(map[string]int)
	counter := make(map[rune]int)
	for i := 1; i < len(polymer); i++ {
		prev := string(polymer[i-1])
		act := string(polymer[i])
		bucket[prev+act]++
	}

	for i := 0; i < limit; i++ {
		newBucket := make(map[string]int)
		for k, v := range bucket {
			a := string(k[0])
			b := string(k[1])
			out := pairs[k]
			newBucket[a+out] += v
			newBucket[out+b] += v
		}

		bucket = newBucket
	}

	counter[rune(polymer[0])]++
	for k, v := range bucket {
		//a := rune(k[0])
		b := rune(k[1])
		//counter[a] += v
		counter[b] += v
	}

	return counter
}
