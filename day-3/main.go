package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Pair struct {
	x int
	y int
}

type CoOrdinates struct {
	i     int
	start int
	end   int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func isSpecialSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func isSafe(x int, y int, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}

func extractNumberCoordinates(x int, y int, engineSchematics []string, n int) (int, int, bool) {
	if !isSafe(x, y, n) || !unicode.IsDigit(rune(engineSchematics[x][y])) {
		return 0, 0, false
	}

	startingX, startingY := x, y
	for {
		if !isSafe(startingX, startingY, n) || !unicode.IsDigit(rune(engineSchematics[startingX][startingY])) {
			break
		}

		startingY--
	}

	endingX, endingY := x, y+1
	for {
		if !isSafe(endingX, endingY, n) || !unicode.IsDigit(rune(engineSchematics[endingX][endingY])) {
			break
		}

		endingY++
	}

	return startingY + 1, endingY, true
}

func findSumPartOne(engineSchematics []string) uint64 {
	directions := []Pair{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	var n int = len(engineSchematics)

	var sum uint64 = 0
	for i := 0; i < len(engineSchematics); i++ {
		for j := 0; j < len(engineSchematics); j++ {
			if isSpecialSymbol(rune(engineSchematics[i][j])) {
				coOrdinates := map[CoOrdinates]struct{}{}
				for _, direction := range directions {
					_x, _y := direction.x+i, direction.y+j
					start, end, ok := extractNumberCoordinates(_x, _y, engineSchematics, n)

					if !ok {
						continue
					}

					coOrdinates[CoOrdinates{i: _x, start: start, end: end}] = struct{}{}
				}

				for coOrdinate := range coOrdinates {
					value, _ := strconv.Atoi(engineSchematics[coOrdinate.i][coOrdinate.start:coOrdinate.end])
					sum += uint64(value)
				}
			}
		}
	}

	return sum
}

func PartOne(path string) uint64 {
	engineSchematics, err := readLines(path)
	if err != nil {
		panic(err)
	}

	return findSumPartOne(engineSchematics)
}

func findSumPartTwo(engineSchematics []string) uint64 {
	directions := []Pair{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}
	var n int = len(engineSchematics)

	var sum uint64 = 0
	for i := 0; i < len(engineSchematics); i++ {
		for j := 0; j < len(engineSchematics); j++ {
			if rune(engineSchematics[i][j]) == '*' {
				coOrdinates := map[CoOrdinates]struct{}{}
				for _, direction := range directions {
					_x, _y := direction.x+i, direction.y+j
					start, end, ok := extractNumberCoordinates(_x, _y, engineSchematics, n)

					if !ok {
						continue
					}

					coOrdinates[CoOrdinates{i: _x, start: start, end: end}] = struct{}{}
				}

				if len(coOrdinates) != 2 {
					continue
				}

				var product uint64 = 1
				for coOrdinate := range coOrdinates {
					value, _ := strconv.Atoi(engineSchematics[coOrdinate.i][coOrdinate.start:coOrdinate.end])
					product *= uint64(value)
				}

				sum += product
			}
		}
	}

	return sum
}

func PartTwo(path string) uint64 {
	engineSchematics, err := readLines(path)
	if err != nil {
		panic(err)
	}

	return findSumPartTwo(engineSchematics)
}

func main() {
	fmt.Println(PartOne("tests/test1.txt"))
	fmt.Println(PartTwo("tests/test2.txt"))
}
