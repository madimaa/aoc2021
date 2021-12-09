package main

import (
	"fmt"
	"os"
	"sort"

	util "github.com/madimaa/aoc2021/util"
	"github.com/madimaa/aoc2021/util/array2d"
)

type pair struct {
	x, y int
}

func main() {
	input := util.OpenFile("input.txt")

	util.Start()
	fmt.Println("Part 1")
	cave := array2d.Create(len(input[0]), len(input))
	for y := 0; y < len(input); y++ {
		line := input[y]
		for x := 0; x < len(line); x++ {
			r := line[x]
			num := util.ConvertToInt(string(r))
			cave.Put(x, y, num)
		}
	}

	lows := part1(cave)
	util.Elapsed()

	util.Start()
	fmt.Println("Part 2")
	part2(cave, lows)
	util.Elapsed()

	os.Exit(0)
}

func part1(cave *array2d.Array2D) []pair {
	maxX, maxY := cave.GetSize()

	lows := make([]pair, 0)
	riskLevelSum := 0

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			act := cave.Get(x, y).(int)
			neighbours := getAdjacents(cave, x, y)
			breaked := false
			for _, p := range neighbours {
				if act >= cave.Get(p.x, p.y).(int) {
					breaked = true
					break
				}
			}

			if breaked {
				continue
			}

			riskLevelSum += 1 + act
			lows = append(lows, pair{x: x, y: y})
		}
	}

	fmt.Println("Result: ", riskLevelSum)
	return lows
}

func part2(cave *array2d.Array2D, lows []pair) {
	basins := make([]int, 0)
	for _, low := range lows {
		basins = append(basins, getBasinSize(cave, low))
	}

	sort.Ints(basins)
	length := len(basins)

	fmt.Println("Result: ", basins[length-1]*basins[length-2]*basins[length-3])
}

func getBasinSize(cave *array2d.Array2D, position pair) int {
	visited := make(map[string]bool)
	haveToVisit := make(map[string]pair)
	haveToVisit[positionToKey(position)] = position

	for len(haveToVisit) > 0 {
		keys := make([]string, 0, len(haveToVisit))
		for k := range haveToVisit {
			keys = append(keys, k)
		}

		for _, key := range keys {
			visited[key] = true
			pos := haveToVisit[key]
			delete(haveToVisit, key)

			neighbours := getAdjacents(cave, pos.x, pos.y)
			for _, p := range neighbours {
				keystring := positionToKey(p)
				if visited[keystring] || cave.Get(p.x, p.y) == 9 {
					continue
				}

				haveToVisit[keystring] = p
			}
		}
	}

	return len(visited)
}

func positionToKey(position pair) string {
	return fmt.Sprintf("%d %d", position.x, position.y)
}

func getAdjacents(a2d *array2d.Array2D, x, y int) []pair {
	xMax, yMax := a2d.GetSize()
	res := make([]pair, 0)

	if x > 0 {
		res = append(res, pair{x: x - 1, y: y})
	}

	if y > 0 {
		res = append(res, pair{x: x, y: y - 1})
	}

	if x < xMax-1 {
		res = append(res, pair{x: x + 1, y: y})
	}

	if y < yMax-1 {
		res = append(res, pair{x: x, y: y + 1})
	}

	return res
}
