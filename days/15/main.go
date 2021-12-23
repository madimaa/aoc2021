package main

import (
	"container/heap"
	"fmt"
	"os"

	util "github.com/madimaa/aoc2021/util"
	"github.com/madimaa/aoc2021/util/array2d"
)

type heuristic func(p *Point) int

func (p *Point) toString() string {
	return fmt.Sprintf("%d %d", p.X, p.Y)
}

func (p *Point) getNeighbours(a2d *array2d.Array2D) []*Point {
	xMax, yMax := a2d.GetSize()
	x, y := p.X, p.Y
	res := make([]*Point, 0)

	if x > 0 {
		res = append(res, a2d.Get(x-1, y).(*Point))
	}

	if y > 0 {
		res = append(res, a2d.Get(x, y-1).(*Point))
	}

	if x < xMax-1 {
		res = append(res, a2d.Get(x+1, y).(*Point))
	}

	if y < yMax-1 {
		res = append(res, a2d.Get(x, y+1).(*Point))
	}

	return res
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
	a2d := array2d.Create(len(input[0]), len(input))
	endX, endY := a2d.GetSize()
	endX--
	endY--

	for y, line := range input {
		for x, r := range line {
			cost := util.ConvertToInt(string(r))
			point := &Point{X: x, Y: y, Cost: cost, Index: 0}
			a2d.Put(x, y, point)
		}
	}

	start := a2d.Get(0, 0).(*Point)
	end := a2d.Get(endX, endY).(*Point)
	h := func(p *Point) int {
		distance := endX - p.X + endY - p.Y
		return distance + p.Cost //*(distance*p.Cost)
	}

	list := aStar(start, end, h, a2d)
	totalRisk := 0
	for _, item := range list {
		//fmt.Println("X: ", item.X, " Y: ", item.Y, " Cost: ", item.Cost)
		totalRisk += item.Cost
	}

	totalRisk -= start.Cost

	fmt.Println("Result: ", totalRisk)
}

func part2(input []string) {
	a2d := array2d.Create(len(input[0])*5, len(input)*5)
	endX, endY := a2d.GetSize()
	endX--
	endY--

	for i := 0; i < 5; i++ {
		yOffset := len(input[0]) * i
		for j := 0; j < 5; j++ {
			xOffset := len(input) * j
			for y, line := range input {
				for x, r := range line {
					cost := util.ConvertToInt(string(r))
					if i != 0 || j != 0 {
						cost += j + i
						if cost > 9 {
							cost = cost % 9
						}
					}
					point := &Point{X: x + xOffset, Y: y + yOffset, Cost: cost, Index: 0}
					a2d.Put(x+xOffset, y+yOffset, point)
				}
			}
		}
	}

	//for y := 0; y <= endY; y++ {
	//	for x := 0; x <= endX; x++ {
	//		fmt.Print(a2d.Get(x, y).(*Point).Cost)
	//	}
	//	fmt.Println()
	//}

	start := a2d.Get(0, 0).(*Point)
	end := a2d.Get(endX, endY).(*Point)
	h := func(p *Point) int {
		distance := endX - p.X + endY - p.Y
		return distance + p.Cost
	}

	list := aStar(start, end, h, a2d)
	totalRisk := 0
	for _, item := range list {
		//fmt.Println("X: ", item.X, " Y: ", item.Y, " Cost: ", item.Cost)
		totalRisk += item.Cost
	}

	totalRisk -= start.Cost

	fmt.Println("Result: ", totalRisk)
}

func d(current, neighbour *Point) int {
	return current.Cost + neighbour.Cost // * (current.Cost + neighbour.Cost)
}

func reconstructPath(cameFrom map[string]*Point, current *Point) []*Point {
	totalPath := make([]*Point, 0)
	totalPath = append(totalPath, current)
	for {
		if val, ok := cameFrom[current.toString()]; ok {
			totalPath = append(totalPath, val)
			current = val
		} else {
			break
		}
	}

	return totalPath
}

func aStar(start, goal *Point, h heuristic, a2d *array2d.Array2D) []*Point {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	start.fScore = h(start)
	openSet.Push(start)
	//openSet = append(openSet, start)

	cameFrom := make(map[string]*Point)

	//it is easier to think about 0 as infinity, so 1 is cheaper than 0.
	gScore := make(map[*Point]int)
	gScore[start] = 0

	//fScore := make(map[*Point]int)
	//fScore[start] = h(start)

	for len(openSet) > 0 {
		current := heap.Pop(&openSet).(*Point)

		if current.toString() == goal.toString() {
			return reconstructPath(cameFrom, current)
		}

		neighbours := current.getNeighbours(a2d)
		for _, neighbour := range neighbours {
			tentativeGScore := gScore[current] + d(current, neighbour)
			if _, ok := gScore[neighbour]; !ok || tentativeGScore < gScore[neighbour] {
				cameFrom[neighbour.toString()] = current
				gScore[neighbour] = tentativeGScore
				neighbour.fScore = tentativeGScore + h(neighbour)
				//fScore[neighbour] = tentativeGScore + h(neighbour)
				contains := false
				for _, item := range openSet {
					if item.toString() == neighbour.toString() {
						contains = true
						break
					}
				}

				if !contains {
					openSet.Push(neighbour)
				}
			}
		}
	}

	panic("Something went wrong.")
}
