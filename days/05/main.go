package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
	"github.com/madimaa/aoc2021/util/array2d"
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
	array := array2d.Create(1000, 1000) //by checking the ipnut file, every number is bellow 1000
	overlap := 0
	for _, line := range input {
		x1y1 := strings.Split(line, " -> ")[0]
		x2y2 := strings.Split(line, " -> ")[1]
		x1 := util.ConvertToInt(strings.Split(x1y1, ",")[0])
		y1 := util.ConvertToInt(strings.Split(x1y1, ",")[1])
		x2 := util.ConvertToInt(strings.Split(x2y2, ",")[0])
		y2 := util.ConvertToInt(strings.Split(x2y2, ",")[1])
		if x1 == x2 || y1 == y2 {
			if y1 > y2 {
				helper := y1
				y1 = y2
				y2 = helper
			}

			if x1 > x2 {
				helper := x1
				x1 = x2
				x2 = helper
			}

			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					if array.Get(x, y) == nil {
						array.Put(x, y, 1)
					} else {
						value := array.Get(x, y).(int)
						if value == 1 {
							overlap++
							array.Put(x, y, value+1)
						}
					}
				}
			}
		}
	}

	fmt.Println("Result: ", overlap)
}

func part2(input []string) {
	array := array2d.Create(1000, 1000) //by checking the ipnut file, every number is bellow 1000
	overlap := 0
	for _, line := range input {
		x1y1 := strings.Split(line, " -> ")[0]
		x2y2 := strings.Split(line, " -> ")[1]
		x1 := util.ConvertToInt(strings.Split(x1y1, ",")[0])
		y1 := util.ConvertToInt(strings.Split(x1y1, ",")[1])
		x2 := util.ConvertToInt(strings.Split(x2y2, ",")[0])
		y2 := util.ConvertToInt(strings.Split(x2y2, ",")[1])
		if x1 == x2 {
			if y1 > y2 {
				helper := y1
				y1 = y2
				y2 = helper
			}

			for y := y1; y <= y2; y++ {
				if array.Get(x1, y) == nil {
					array.Put(x1, y, 1)
				} else {
					value := array.Get(x1, y).(int)
					if value == 1 {
						overlap++
						array.Put(x1, y, value+1)
					}
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				helper := x1
				x1 = x2
				x2 = helper
			}

			for x := x1; x <= x2; x++ {
				if array.Get(x, y1) == nil {
					array.Put(x, y1, 1)
				} else {
					value := array.Get(x, y1).(int)
					if value == 1 {
						overlap++
						array.Put(x, y1, value+1)
					}
				}
			}
		} else {
			for {
				if array.Get(x1, y1) == nil {
					array.Put(x1, y1, 1)
				} else {
					value := array.Get(x1, y1).(int)
					if value == 1 {
						overlap++
						array.Put(x1, y1, value+1)
					}
				}

				if x1 == x2 && y1 == y2 {
					break
				}

				if x1 > x2 {
					x1--
				} else {
					x1++
				}

				if y1 > y2 {
					y1--
				} else {
					y1++
				}
			}
		}
	}

	fmt.Println("Result: ", overlap)
}
