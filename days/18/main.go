package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
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
	actual := input[0]
	for i := 1; i < len(input); i++ {
		actual = fmt.Sprintf("[%s,%s]", actual, input[i])
		actual = doReduction(actual)
	}

	magnitude := calculateMagnitude(actual)
	fmt.Println("Result: ", magnitude)
}

func part2(input []string) {
	largest := 0
	for j := 0; j < len(input); j++ {
		for i := 0; i < len(input); i++ {
			if i == j {
				continue
			}
			actual := input[j]
			actual = fmt.Sprintf("[%s,%s]", actual, input[i])
			actual = doReduction(actual)

			magnitude := calculateMagnitude(actual)
			m := util.ConvertToInt(magnitude)
			if m > largest {
				largest = m
			}
		}

	}

	fmt.Println("Result: ", largest)
}

func doReduction(input string) string {
	result := input
	level := 0
	indexof := 0
loop:
	for i, r := range input {
		switch r {
		case '[':
			level++
			if level == 5 {
				indexof = i
				break loop
			}
		case ']':
			level--
		}
	}

	if level == 5 {
		return doReduction(doExplode(input, indexof))
	}

	r, _ := regexp.Compile(`[0-9]{2}`)
	location := r.FindStringIndex(input)

	if location != nil {
		return doReduction(doSplit(input, location[0]))
	}

	return result
}

func doExplode(input string, indexStart int) string {
	shouldReplace := ""
	indexEnd := 0
	for i := indexStart; i < len(input); i++ {
		shouldReplace += string(input[i])
		if input[i] == ']' {
			indexEnd = i
			break
		}
	}

	indexOfLeftStart, indexOfLeftEnd, indexOfRightStart, indexOfRightEnd := 0, 0, 0, 0
	for i := indexStart; i >= 0; i-- {
		if input[i] != ',' && input[i] != '[' && input[i] != ']' {
			indexOfLeftEnd = i
			indexOfLeftStart = i
			break
		}
	}

	for i := indexOfLeftEnd; i >= 0; i-- {
		if input[i] == ',' || input[i] == '[' || input[i] == ']' {
			break
		}
		indexOfLeftStart = i
	}

	for i := indexStart + len(shouldReplace); i < len(input); i++ {
		if input[i] != ',' && input[i] != '[' && input[i] != ']' {
			indexOfRightStart = i
			indexOfRightEnd = i
			break
		}
	}

	for i := indexOfRightStart; i < len(input); i++ {
		if input[i] == ',' || input[i] == '[' || input[i] == ']' {
			break
		}
		indexOfRightEnd = i
	}

	result := ""
	trimmed := shouldReplace[1 : len(shouldReplace)-1]
	leftNum := util.ConvertToInt(strings.Split(trimmed, ",")[0])
	rightNum := util.ConvertToInt(strings.Split(trimmed, ",")[1])
	if indexOfLeftEnd > 0 && indexOfRightStart > 0 {
		left := util.ConvertToInt(input[indexOfLeftStart:indexOfLeftEnd+1]) + leftNum
		right := util.ConvertToInt(input[indexOfRightStart:indexOfRightEnd+1]) + rightNum
		result += input[:indexOfLeftStart]
		result += fmt.Sprint(left)
		result += input[indexOfLeftEnd+1 : indexStart]
		result += "0"
		result += input[indexEnd+1 : indexOfRightStart]
		result += fmt.Sprint(right)
		result += input[indexOfRightEnd+1:]
	} else if indexOfLeftEnd > 0 {
		left := util.ConvertToInt(input[indexOfLeftStart:indexOfLeftEnd+1]) + leftNum
		result = fmt.Sprintf("%s%d%s%d%s", input[:indexOfLeftStart], left, input[indexOfLeftEnd+1:indexStart], 0, input[indexEnd+1:])
	} else if indexOfRightStart > 0 {
		right := util.ConvertToInt(input[indexOfRightStart:indexOfRightEnd+1]) + rightNum
		result = fmt.Sprintf("%s%d%s%d%s", input[:indexStart], 0, input[indexEnd+1:indexOfRightStart], right, input[indexOfRightEnd+1:])
	} else {
		panic("Something went wrong!")
	}

	return result
}

func doSplit(input string, index int) string {
	r, _ := regexp.Compile(`[0-9]{2}`)
	str := r.FindString(input)
	num := util.ConvertToInt(str)
	left := math.Floor(float64(num) / 2)
	right := math.Ceil(float64(num) / 2)

	result := fmt.Sprintf("%s[%d,%d]%s", input[:index], int(left), int(right), input[index+2:])

	return result
}

func calculateMagnitude(input string) string {
	r, _ := regexp.Compile(`\[[0-9]+,[0-9]+\]`)
	for {
		if !r.MatchString(input) {
			break
		}

		result := r.FindString(input)
		trimmed := result[1 : len(result)-1]
		left := util.ConvertToInt(strings.Split(trimmed, ",")[0]) * 3
		right := util.ConvertToInt(strings.Split(trimmed, ",")[1]) * 2
		input = strings.Replace(input, result, fmt.Sprintf("%d", left+right), 1)
	}

	return input
}
