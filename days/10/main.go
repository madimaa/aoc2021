package main

import (
	"fmt"
	"os"
	"sort"

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
	errorPoints := 0

	for _, line := range input {
		opens := make([]rune, 0)
		corrupted := false
		corruptedRune := '0'
	loop:
		for _, r := range line {
			switch r {
			case '<', '{', '(', '[':
				opens = append(opens, r)
			case '>':
				if opens[len(opens)-1] != '<' {
					corrupted = true
					corruptedRune = r
					break loop
				}
				opens = opens[:len(opens)-1]
			case '}':
				if opens[len(opens)-1] != '{' {
					corrupted = true
					corruptedRune = r
					break loop
				}
				opens = opens[:len(opens)-1]
			case ')':
				if opens[len(opens)-1] != '(' {
					corrupted = true
					corruptedRune = r
					break loop
				}
				opens = opens[:len(opens)-1]
			case ']':
				if opens[len(opens)-1] != '[' {
					corrupted = true
					corruptedRune = r
					break loop
				}
				opens = opens[:len(opens)-1]
			}
		}

		if corrupted {
			switch corruptedRune {
			case '>':
				errorPoints += 25137
			case '}':
				errorPoints += 1197
			case ')':
				errorPoints += 3
			case ']':
				errorPoints += 57
			}
		}
	}

	fmt.Println("Result: ", errorPoints)
}

func part2(input []string) {
	autocompletePoints := make([]int, 0)

loop:
	for _, line := range input {
		opens := make([]rune, 0)
		for _, r := range line {
			switch r {
			case '<', '{', '(', '[':
				opens = append(opens, r)
			case '>':
				if opens[len(opens)-1] != '<' {
					continue loop
				}
				opens = opens[:len(opens)-1]
			case '}':
				if opens[len(opens)-1] != '{' {
					continue loop
				}
				opens = opens[:len(opens)-1]
			case ')':
				if opens[len(opens)-1] != '(' {
					continue loop
				}
				opens = opens[:len(opens)-1]
			case ']':
				if opens[len(opens)-1] != '[' {
					continue loop
				}
				opens = opens[:len(opens)-1]
			}
		}

		points := 0
		for i := len(opens) - 1; i >= 0; i-- {
			points *= 5
			r := opens[i]
			switch r {
			case '<':
				points += 4
			case '{':
				points += 3
			case '(':
				points += 1
			case '[':
				points += 2
			}
		}

		autocompletePoints = append(autocompletePoints, points)
	}

	sort.Ints(autocompletePoints)

	fmt.Println("Result: ", autocompletePoints[len(autocompletePoints)/2])
}
