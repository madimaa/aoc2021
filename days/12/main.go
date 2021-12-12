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
	connections := getConnections(input)
	paths := buildPath(connections, []string{"start"})

	fmt.Println("Result: ", len(paths))
}

func part2(input []string) {
	connections := getConnections(input)
	paths := buildPathPart2(connections, []string{"start"})

	fmt.Println("Result: ", len(paths))
}

func getConnections(input []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, "-")
		if _, ok := connections[split[0]]; !ok {
			connections[split[0]] = make([]string, 0)
		}

		connections[split[0]] = append(connections[split[0]], split[1])

		if _, ok := connections[split[1]]; !ok {
			connections[split[1]] = make([]string, 0)
		}

		connections[split[1]] = append(connections[split[1]], split[0])
	}

	return connections
}

func buildPath(connections map[string][]string, actualPath []string) [][]string {
	paths := make([][]string, 0)
	lastElement := actualPath[len(actualPath)-1]
	for _, val := range connections[lastElement] {
		path := make([]string, 0)
		if val == "end" {
			path = append(path, actualPath...)
			path = append(path, val)
			paths = append(paths, path)
			continue
		}

		if util.ContainsStr(actualPath, val) {
			if strings.ToUpper(val) == val {
				path = append(path, actualPath...)
				path = append(path, val)
				paths = append(paths, buildPath(connections, path)...)
			}
		} else {
			path = append(path, actualPath...)
			path = append(path, val)
			paths = append(paths, buildPath(connections, path)...)
		}

	}

	return paths
}

func buildPathPart2(connections map[string][]string, actualPath []string) [][]string {
	paths := make([][]string, 0)
	lastElement := actualPath[len(actualPath)-1]
	for _, val := range connections[lastElement] {
		path := make([]string, 0)
		if val == "end" {
			path = append(path, actualPath...)
			path = append(path, val)
			paths = append(paths, path)
			continue
		}

		if val == "start" {
			continue
		}

		if util.ContainsStr(actualPath, val) {
			if strings.ToUpper(val) == val {
				path = append(path, actualPath...)
				path = append(path, val)
				paths = append(paths, buildPathPart2(connections, path)...)
			} else {
				if !isThereAny2Lower(actualPath, val) {
					path = append(path, actualPath...)
					path = append(path, val)
					paths = append(paths, buildPathPart2(connections, path)...)
				}
			}
		} else {
			path = append(path, actualPath...)
			path = append(path, val)
			paths = append(paths, buildPathPart2(connections, path)...)
		}

	}

	return paths
}

func isThereAny2Lower(actualPath []string, val string) bool {
	fakeCounter := make(map[string]bool)
	for _, item := range actualPath {
		if item == strings.ToUpper(item) {
			continue
		}

		if fakeCounter[item] {
			return true
		} else {
			fakeCounter[item] = true
		}
	}

	return false
}
