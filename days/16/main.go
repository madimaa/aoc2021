package main

import (
	"fmt"
	"os"
	"strconv"

	util "github.com/madimaa/aoc2021/util"
)

var _version int64

const (
	sumPocket         int64 = 0
	productPocket     int64 = 1
	minimumPocket     int64 = 2
	maximumPocket     int64 = 3
	greaterThanPocket int64 = 5
	lessThanPocket    int64 = 6
	equalToPocket     int64 = 7
)

func main() {
	input := util.OpenFile("input.txt")

	util.Start()
	solve(input[0])
	util.Elapsed()

	os.Exit(0)
}

func solve(input string) {
	binaryString := ""
	for _, r := range input {
		val := convertHexToBinaryString(string(r))
		binaryString += val
	}

	_version = 0
	_, result := handlePacket(binaryString, 0)

	fmt.Println("Part1: ", _version)
	fmt.Println("Part2: ", result)
}

func convertHexToBinaryString(in string) string {
	if val, err := strconv.ParseInt(in, 16, 64); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%04b", val)
	}
}

func convertBinaryStringToDecimal(in string) int64 {
	if val, err := strconv.ParseInt(in, 2, 64); err != nil {
		panic(err)
	} else {
		return val
	}
}

func handlePacket(binaryString string, lastHandledBit int) (int, int) {
	//fmt.Println("Binary string: ", binaryString)
	version := convertBinaryStringToDecimal(binaryString[lastHandledBit : lastHandledBit+3])
	lastHandledBit += 3
	//fmt.Println("Version: ", version)
	_version += version
	typeId := convertBinaryStringToDecimal(binaryString[lastHandledBit : lastHandledBit+3])
	lastHandledBit += 3
	//fmt.Println("Type: ", typeId)
	switch typeId {
	case 4:
		//Not An Opretator
		numberString := ""
		for {
			bits := binaryString[lastHandledBit : lastHandledBit+5]
			numberString += bits[1:]
			lastHandledBit += 5
			if bits[0] == '0' {
				//last group
				break
			}
		}
		number := convertBinaryStringToDecimal(numberString)
		//fmt.Println(number)
		return lastHandledBit, int(number)
	default:
		//Operator
		lengthTypeId := binaryString[lastHandledBit : lastHandledBit+1]
		lastHandledBit += 1
		responses := make([]int, 0)
		//fmt.Println(lengthTypeId)
		switch lengthTypeId {
		case "0":
			lengthOfSubPackets := convertBinaryStringToDecimal(binaryString[lastHandledBit : lastHandledBit+15])
			lastHandledBit += 15
			endOfSubPackets := int(lengthOfSubPackets) + lastHandledBit
			for {
				res := 0
				lastHandledBit, res = handlePacket(binaryString, lastHandledBit)
				responses = append(responses, res)

				if lastHandledBit >= endOfSubPackets {
					break
				}
			}
		case "1":
			numberOfSubPackets := convertBinaryStringToDecimal(binaryString[lastHandledBit : lastHandledBit+11])
			lastHandledBit += 11
			restOfTheString := binaryString[lastHandledBit:]
			//fmt.Println(numberOfSubPackets)
			lengthNow := len(restOfTheString)
			for i := 0; i < int(numberOfSubPackets); i++ {
				res := 0
				lastHandledBit, res = handlePacket(binaryString, lastHandledBit)
				responses = append(responses, res)
			}
			lengthAfter := len(restOfTheString)
			lastHandledBit += lengthNow - lengthAfter
		default:
			panic("¯\\_(ツ)_/¯")
		}

		switch typeId {
		case sumPocket:
			sum := 0
			for _, num := range responses {
				sum += num
			}
			return lastHandledBit, sum
		case productPocket:
			prod := 1
			for _, num := range responses {
				prod *= num
			}
			return lastHandledBit, prod
		case minimumPocket:
			min := responses[0]
			for i := 1; i < len(responses); i++ {
				if responses[i] < min {
					min = responses[i]
				}
			}
			return lastHandledBit, min
		case maximumPocket:
			max := responses[0]
			for i := 1; i < len(responses); i++ {
				if responses[i] > max {
					max = responses[i]
				}
			}
			return lastHandledBit, max
		case greaterThanPocket:
			if responses[0] > responses[1] {
				return lastHandledBit, 1
			} else {
				return lastHandledBit, 0
			}
		case lessThanPocket:
			if responses[0] < responses[1] {
				return lastHandledBit, 1
			} else {
				return lastHandledBit, 0
			}
		case equalToPocket:
			if responses[0] == responses[1] {
				return lastHandledBit, 1
			} else {
				return lastHandledBit, 0
			}
		default:
			panic("Something went wrong!")
		}
	}
}
