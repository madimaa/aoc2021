package main

import (
	"fmt"
	"os"
	"strings"

	util "github.com/madimaa/aoc2021/util"
)

type Vector struct {
	x, y int
}

type Probe struct {
	x, y     int
	velocity *Vector
}

func (p *Probe) step() {
	p.x += p.velocity.x
	p.y += p.velocity.y
	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}

	p.velocity.y--
}

func (p *Probe) isInTargetArea(xMin, xMax, yMin, yMax int) bool {
	if p.x >= xMin && p.x <= xMax && p.y >= yMin && p.y <= yMax {
		return true
	} else {
		return false
	}
}

func main() {
	input := util.OpenFile("input.txt")

	util.Start()
	solve(input[0])
	util.Elapsed()

	os.Exit(0)
}

func solve(input string) {
	xMin, xMax, yMin, yMax := getTargetArea(input)
	highest := 0
	successful := 0
	for x := -200; x < 200; x++ {
		for y := -200; y < 200; y++ {
			probe := &Probe{x: 0, y: 0, velocity: &Vector{x: x, y: y}}
			actualHighest := 0
			for i := 0; i < 500; i++ {
				probe.step()
				if probe.y > actualHighest {
					actualHighest = probe.y
				}

				if probe.isInTargetArea(xMin, xMax, yMin, yMax) {
					if highest < actualHighest {
						highest = actualHighest
					}
					successful++
					break
				}

				if probe.velocity.y < 0 && probe.y < yMin {
					break
				}

				if probe.velocity.x == 0 && (probe.x > xMax || probe.x < xMin) {
					break
				}
			}
		}
	}

	fmt.Println("Part 1: ", highest)
	fmt.Println("Part 2: ", successful)
}

func getTargetArea(input string) (int, int, int, int) {
	input = strings.TrimPrefix(input, "target area: ")
	xString := strings.TrimPrefix(strings.Split(input, ", ")[0], "x=")
	yString := strings.TrimPrefix(strings.Split(input, ", ")[1], "y=")
	xMin, xMax := getMinMaxAsInt(xString)
	yMin, yMax := getMinMaxAsInt(yString)
	return xMin, xMax, yMin, yMax
}

func getMinMaxAsInt(input string) (int, int) {
	min := util.ConvertToInt(strings.Split(input, "..")[0])
	max := util.ConvertToInt(strings.Split(input, "..")[1])
	return min, max
}
