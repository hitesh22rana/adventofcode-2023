package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	id       uint64
	red      uint64
	green    uint64
	blue     uint64
	possible bool
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

func extractRecordPartOne(rawRecord string) Record {
	var record Record = Record{
		id:       0,
		red:      0,
		green:    0,
		blue:     0,
		possible: true,
	}

	extractedrawData := strings.Split(rawRecord, ":")
	id, _ := strconv.Atoi(strings.Split(extractedrawData[0], " ")[1])
	record.id = uint64(id)

	gameRecords := strings.Split(extractedrawData[1], ";")

	for _, gameRecord := range gameRecords {
		gameRecordsData := strings.Split(gameRecord, ",")

		for _, gameRecordData := range gameRecordsData {
			data := strings.Split(gameRecordData, " ")

			red, green, blue := 0, 0, 0
			color := data[2]
			value, _ := strconv.Atoi(data[1])
			switch color {
			case "red":
				record.red += uint64(value)
				red = value
			case "green":
				record.green += uint64(value)
				green = value
			case "blue":
				record.blue += uint64(value)
				blue = value
			}

			record.possible = record.possible && (red <= 12 && green <= 13 && blue <= 14)
		}
	}

	return record
}

func PartOne(path string) uint64 {
	records, err := readLines(path)
	if err != nil {
		panic(err)
	}

	var sum uint64 = 0
	for _, rawRecord := range records {
		record := extractRecordPartOne(rawRecord)
		if record.possible {
			sum += record.id
		}
	}

	return sum
}

func extractRecordPartTwo(rawRecord string) Record {
	var record Record

	extractedrawData := strings.Split(rawRecord, ":")
	id, _ := strconv.Atoi(strings.Split(extractedrawData[0], " ")[1])
	record.id = uint64(id)

	gameRecords := strings.Split(extractedrawData[1], ";")

	for _, gameRecord := range gameRecords {
		gameRecordsData := strings.Split(gameRecord, ",")

		for _, gameRecordData := range gameRecordsData {
			data := strings.Split(gameRecordData, " ")

			color := data[2]
			value, _ := strconv.Atoi(data[1])
			switch color {
			case "red":
				record.red = uint64(math.Max(float64(value), float64(record.red)))
			case "green":
				record.green = uint64(math.Max(float64(value), float64(record.green)))
			case "blue":
				record.blue = uint64(math.Max(float64(value), float64(record.blue)))
			}
		}
	}

	return record
}

func (record Record) minimumCubesMultipliedTogether() uint64 {
	return record.red * record.green * record.blue
}

func PartTwo(path string) uint64 {
	records, err := readLines(path)
	if err != nil {
		panic(err)
	}

	var sum uint64 = 0
	for _, rawRecord := range records {
		record := extractRecordPartTwo(rawRecord)
		sum += record.minimumCubesMultipliedTogether()
	}

	return sum
}

func main() {
	fmt.Println(PartOne("tests/test1.txt"))
	fmt.Println(PartTwo("tests/test2.txt"))
}
