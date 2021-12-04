package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

type marked struct {
	data []int
}

type won struct {
	data bool
}

type board struct {
	data          [][]int
	markedNumbers *marked
	bingo         *won
}

func (b *board) isBingo(input int) bool {
	bingo := false

	for i := 1; i < len(b.data); i++ {
		line := b.data[i]
		for j := 1; j < len(line); j++ {
			if input == line[j] {
				b.data[i][0]++
				b.data[0][j]++
				b.markedNumbers.data = append(b.markedNumbers.data, line[j])
			}
		}
	}

	target := len(b.data) - 1
	for i, line := range b.data {
		for j, column := range line {
			if j == 0 || i == 0 {
				if target == column {
					bingo = true
					b.bingo.data = true
					break
				}
			}
		}
	}

	return bingo
}

func (b *board) sumRemaining() int {
	sum := 0
	for i := 1; i < len(b.data); i++ {
		line := b.data[i]
		for j := 1; j < len(line); j++ {
			if !util.ContainsInt(b.markedNumbers.data, line[j]) {
				sum += line[j]
			}
		}
	}

	return sum
}

func main() {
	input := util.OpenFile("input.txt")

	util.Start()
	fmt.Println("Part 1")
	solve(input)
	util.Elapsed()

	os.Exit(0)
}

func solve(input []string) {
	numbers := getDrawnNumbers(input[0])
	boards := getBoards(input[2:])
	var part1 int
	var part2 int
	first := true
	for _, number := range numbers {
		for _, b := range boards {
			if b.bingo.data {
				continue
			}

			if b.isBingo(number) {
				result := b.sumRemaining() * number
				if first {
					part1 = result
					first = false
				}

				part2 = result
			}
		}
	}

	fmt.Println("Result of part 1: ", part1)
	fmt.Println("Result of part 2: ", part2)
}

func getDrawnNumbers(input string) []int {
	result := make([]int, 0)
	for _, v := range strings.Split(input, ",") {
		num, err := strconv.Atoi(v)
		util.PanicOnError(err)
		result = append(result, num)
	}

	return result
}

func getBoards(input []string) []board {
	boards := make([]board, 0)

	board_ := board{data: make([][]int, 0), markedNumbers: &marked{data: make([]int, 0)}, bingo: &won{}}
	first := true
	re := regexp.MustCompile("[[:space:]]+")
	for _, line := range input {
		if len(line) == 0 {
			boards = append(boards, board_)
			board_ = board{data: make([][]int, 0), markedNumbers: &marked{data: make([]int, 0)}, bingo: &won{}}
			first = true
			continue
		}

		numbers := make([]int, 0)
		numbers = append(numbers, 0)
		for _, v := range re.Split(strings.TrimSpace(line), -1) {
			num, err := strconv.Atoi(v)
			util.PanicOnError(err)
			numbers = append(numbers, num)
		}

		if first {
			board_.data = append(board_.data, make([]int, len(numbers)))
			first = false
		}
		board_.data = append(board_.data, numbers)
	}

	return boards
}
