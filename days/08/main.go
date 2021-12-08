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
	simples := 0
	for _, line := range input {
		outputs := strings.Split(strings.Split(line, " | ")[1], " ")
		for _, output := range outputs {
			length := len(output)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				simples++
			}
		}
	}

	fmt.Println("Result: ", simples)
}

func part2(input []string) {
	sum := 0
	for _, line := range input {
		outputs := strings.Split(strings.Split(line, " | ")[1], " ")
		inputs := strings.Split(strings.Split(line, " | ")[0], " ")
		zero := ""
		one := ""
		two := ""
		three := ""
		four := ""
		five := ""
		six := ""
		seven := ""
		eight := ""
		nine := ""

		for _, val := range inputs {
			switch len(val) {
			case 2:
				one = val
			case 3:
				seven = val
			case 4:
				four = val
			case 7:
				eight = val
			}
		}

		for _, val := range inputs {
			if len(val) == 5 {
				if isMatch(val, one) {
					three = val
				}
			}
		}

		for _, val := range inputs {
			if len(val) == 6 {
				if isMatch(val, four) {
					nine = val
				}
			}
		}

		dddd, bb := getSegments(one, four, three)

		for _, val := range inputs {
			if val == nine {
				continue
			}

			if len(val) == 6 {
				if strings.ContainsRune(val, dddd) {
					six = val
				} else {
					zero = val
				}
			}
		}

		for _, val := range inputs {
			if val == three {
				continue
			}

			if len(val) == 5 {
				if strings.ContainsRune(val, bb) {
					five = val
				} else {
					two = val
				}
			}
		}

		num := ""
		for _, output := range outputs {
			if isMatch(output, eight) {
				num += "8"
			} else if isMatch(output, zero) {
				num += "0"
			} else if isMatch(output, nine) {
				num += "9"
			} else if isMatch(output, six) {
				num += "6"
			} else if isMatch(output, two) {
				num += "2"
			} else if isMatch(output, three) {
				num += "3"
			} else if isMatch(output, four) {
				num += "4"
			} else if isMatch(output, five) {
				num += "5"
			} else if isMatch(output, seven) {
				num += "7"
			} else if isMatch(output, one) {
				num += "1"
			} else {
				panic("Something went wrong!")
			}
		}

		sum += util.ConvertToInt(num)
	}

	fmt.Println("Result: ", sum)
}

func isMatch(val, target string) bool {
	for _, r := range target {
		maybe := false
		for _, x := range val {
			if r == x {
				maybe = true
			}
		}

		if !maybe {
			return false
		}
	}

	return true
}

func getSegments(one, four, three string) (rune, rune) {
	less := ""
	for _, r1 := range four {
		contains := false

		for _, r2 := range one {
			if r1 == r2 {
				contains = true
			}
		}

		if !contains {
			less += string(r1)
		}
	}

	var dddd rune
	for _, r1 := range less {
		for _, r2 := range three {
			if r1 == r2 {
				dddd = r1
			}
		}
	}

	var bb rune
	for _, r := range less {
		if r != dddd {
			bb = r
		}
	}

	return dddd, bb
}
