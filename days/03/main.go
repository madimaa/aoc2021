package main

import (
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2021/util"
)

type rating int8

const (
	oxygen rating = 0
	co2    rating = 1
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
	rate := make([]int, len(input[0]))
	for _, line := range input {
		for i, char := range line {
			if char == '0' {
				rate[i]--
			} else {
				rate[i]++
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, val := range rate {
		if val > 0 {
			gammaRate += "1"
			epsilonRate += "0"
		} else if val < 0 {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			panic("Instruction unclear... :D")
		}
	}

	gamma := convert(gammaRate)
	epsilon := convert(epsilonRate)

	fmt.Println("Result: ", epsilon*gamma)
}

func part2(input []string) {
	getlast := func(type_ rating) string {
		index := 0
		res := generateRate(input, index, type_)
		for len(res) > 1 {
			index++
			res = generateRate(res, index, type_)
		}

		return res[0]
	}

	oxygen := convert(getlast(oxygen))
	co2 := convert(getlast(co2))

	fmt.Println("Result: ", oxygen*co2)
}

func generateRate(input []string, index int, type_ rating) []string {
	rate := 0
	for _, line := range input {
		if line[index] == '0' {
			rate--
		} else {
			rate++
		}
	}

	var criteria rune
	switch type_ {
	case oxygen:
		if rate >= 0 {
			criteria = '1'
		} else {
			criteria = '0'
		}
	case co2:
		if rate >= 0 {
			criteria = '0'
		} else {
			criteria = '1'
		}
	}

	result := make([]string, 0)
	for _, line := range input {
		if line[index] == byte(criteria) {
			result = append(result, line)
		}
	}

	return result
}

func convert(in string) int64 {
	if val, err := strconv.ParseInt(in, 2, 64); err != nil {
		panic(err)
	} else {
		return val
	}
}
