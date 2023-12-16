package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CardGameData struct {
	id  int
	sum uint64
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

func stringsToInt(vals []string) []int {
	var nums []int
	for _, val := range vals {
		if val == "" {
			continue
		}

		num, _ := strconv.Atoi(strings.TrimSpace(val))
		nums = append(nums, num)
	}
	return nums
}

func extractCardGameNumbers(cardGame string) ([]int, []int) {
	rawGameData := strings.Split(cardGame, ": ")

	numbers := strings.Split(rawGameData[1], " | ")
	return stringsToInt(strings.Split(strings.TrimSpace(numbers[0]), " ")), stringsToInt(strings.Split(strings.TrimSpace(numbers[1]), " "))
}

func findSum(cardGame string) (uint64, uint64) {
	winningNumbers, numbersHave := extractCardGameNumbers(cardGame)

	pointsMap := map[int]struct{}{}

	for _, numberHave := range numbersHave {
		pointsMap[numberHave] = struct{}{}
	}

	var points uint64 = 0
	for _, winningNumber := range winningNumbers {
		if _, ok := pointsMap[winningNumber]; ok {
			delete(pointsMap, winningNumber)
			points += 1
		}
	}

	return uint64(1 << (points - 1)), points
}

func PartOne(path string) uint64 {
	cardGames, err := readLines(path)
	if err != nil {
		panic(err)
	}

	var sum uint64 = 0
	for _, cardGame := range cardGames {
		product, _ := findSum(cardGame)
		sum += product
	}
	return sum
}

func findSumPartTwo(cardGames []string) uint64 {
	var n int = len(cardGames)

	cardData := make([]CardGameData, n)
	for i, cardGame := range cardGames {
		cardData[i].id = i
		_, cardData[i].sum = findSum(cardGame)
	}

	values := make([]uint64, n)
	var dataQueue []int

	for i := 0; i < n; i++ {
		values[i] = 1
		dataQueue = append(dataQueue, i)
	}

	for len(dataQueue) > 0 {
		index := dataQueue[0]
		dataQueue = dataQueue[1:]
		data := cardData[index]

		for i := 1; i <= int(data.sum); i++ {
			dataQueue = append(dataQueue, data.id+i)
			values[data.id+i]++
		}
	}

	var sum uint64 = 0
	for _, value := range values {
		sum += value
	}

	return sum
}

func PartTwo(path string) uint64 {
	cardGames, err := readLines(path)
	if err != nil {
		panic(err)
	}

	return findSumPartTwo(cardGames)
}

func main() {
	fmt.Println(PartOne("tests/test1.txt"))
	fmt.Println(PartTwo("tests/test2.txt"))
}
