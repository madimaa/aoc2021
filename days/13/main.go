package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
	"github.com/madimaa/aoc2021/util/array2d"
)

type instruction struct {
	axis     rune
	location int
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
	fmt.Println("Result: ", solve(input, "part1"))
}

func part2(input []string) {

	fmt.Println("Result: ", solve(input, "part2"))
}

func getActualSize(arr *array2d.Array2D) (int, int) {
	actX, actY := arr.GetSize()
	maxX, maxY := 0, 0

	for y := 0; y < actY; y++ {
		for x := 0; x < actX; x++ {
			if arr.Get(x, y) != nil {
				if x > maxX {
					maxX = x
				}

				if y > maxY {
					maxY = y
				}
			}
		}
	}

	return maxX + 1, maxY + 1
}

func copyArr(arr *array2d.Array2D, maxX, maxY int) *array2d.Array2D {
	copy_ := array2d.Create(maxX, maxY)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if arr.Get(x, y) != nil {
				copy_.Put(x, y, '#')
			}
		}
	}

	return copy_
}

func solve(input []string, part string) int {
	data := true
	instructions := make([]instruction, 0)
	arr := array2d.Create(1500, 1500)
	for _, line := range input {
		if len(line) == 0 {
			data = false
			continue
		}

		if data {
			split := strings.Split(line, ",")
			x := util.ConvertToInt(split[0])
			y := util.ConvertToInt(split[1])
			arr.Put(x, y, '#')
		} else {
			split := strings.Split(line, "=")
			num := util.ConvertToInt(split[1])
			r := split[0][len(split[0])-1]
			instructions = append(instructions, instruction{axis: rune(r), location: num})
		}
	}

	// It makes it a  bit slower, but its AoC.
	maxX, maxY := getActualSize(arr)
	arr = copyArr(arr, maxX, maxY)
	//visualize(arr)

	for _, inst := range instructions {
		newArr := array2d.Create(maxX, maxY)
		if inst.axis == 'x' {
			mirror := inst.location
			for y := 0; y < maxY; y++ {
				for x := 0; x < mirror; x++ {
					if arr.Get(x, y) != nil {
						newArr.Put(x, y, '#')
					}
				}
			}

			for y := 0; y < maxY; y++ {
				for x := mirror + 1; x < maxX; x++ {
					if arr.Get(x, y) != nil {
						newX := mirror - x + mirror
						newArr.Put(newX, y, '#')
					}
				}
			}
		} else {
			mirror := inst.location
			for y := 0; y < mirror; y++ {
				for x := 0; x < maxX; x++ {
					if arr.Get(x, y) != nil {
						newArr.Put(x, y, '#')
					}
				}
			}

			for y := mirror + 1; y < maxY; y++ {
				for x := 0; x < maxX; x++ {
					if arr.Get(x, y) != nil {
						newY := mirror - y + mirror
						newArr.Put(x, newY, '#')
					}
				}
			}
		}

		maxX, maxY = getActualSize(newArr)
		arr = copyArr(newArr, maxX, maxY)
		if part == "part1" {
			break
		}
	}

	dotcount := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if arr.Get(x, y) != nil {
				dotcount++
			}
		}
	}

	if part == "part2" {
		visualize(arr)
	}

	return dotcount
}

func visualize(arr *array2d.Array2D) {
	fmt.Println("-----------------------------------------------")
	maxX, maxY := arr.GetSize()
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if arr.Get(x, y) != nil {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------------------------------------")
}
