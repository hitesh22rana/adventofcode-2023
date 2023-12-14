package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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

func numberOnlyWithDigits(word string) uint64 {
	length := len(word)

	var str string
	for i := 0; i < length; i++ {
		if unicode.IsDigit(rune(word[i])) {
			str += string(word[i])
			break
		}
	}

	for i := length - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(word[i])) {
			str += string(word[i])
			break
		}
	}

	num, _ := strconv.Atoi(str)
	return uint64(num)
}

func PartOne(path string) uint64 {
	words, err := readLines(path)
	if err != nil {
		panic(err)
	}

	var sum uint64 = 0
	for _, word := range words {
		sum += numberOnlyWithDigits(word)
	}

	return sum
}

func numberWithWordsAndDigits(word string) uint64 {
	length := len(word)

	wordsDigits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	firstDigitIndex := -1
	for i := 0; i < length; i++ {
		if unicode.IsDigit(rune(word[i])) {
			firstDigitIndex = i
			break
		}
	}

	lastDigitIndex := -1
	for i := length - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(word[i])) {
			lastDigitIndex = i
			break
		}
	}

	indxMapFromStart := make(map[int]string)
	for value := range wordsDigits {
		indx := strings.Index(word, value)
		if indx != -1 {
			indxMapFromStart[indx] = value
		}
	}

	firstWordIndx := length
	for indx := range indxMapFromStart {
		if indx < firstWordIndx {
			firstWordIndx = indx
		}
	}

	indxMapFromEnd := make(map[int]string)
	for value := range wordsDigits {
		indx := strings.LastIndex(word, value)
		if indx != -1 {
			indxMapFromEnd[indx] = value
		}
	}

	lastWordIndx := -1
	for indx := range indxMapFromEnd {
		if indx > lastWordIndx {
			lastWordIndx = indx
		}
	}

	var sum int = 0
	if firstWordIndx < firstDigitIndex || firstDigitIndex == -1 {
		sum = wordsDigits[indxMapFromStart[firstWordIndx]] * 10
	} else {
		val, _ := strconv.Atoi(string(word[firstDigitIndex]))
		sum = val * 10
	}

	if lastWordIndx > lastDigitIndex || lastDigitIndex == -1 {
		sum += wordsDigits[indxMapFromEnd[lastWordIndx]]
	} else {
		val, _ := strconv.Atoi(string(word[lastDigitIndex]))
		sum += val
	}

	return uint64(sum)
}

func PartTwo(path string) uint64 {
	words, err := readLines(path)
	if err != nil {
		panic(err)
	}

	var sum uint64 = 0
	for _, word := range words {
		sum += numberWithWordsAndDigits(word)
	}

	return sum
}

func main() {
	fmt.Println(PartOne("tests/test1.txt"))
	fmt.Println(PartTwo("tests/test2.txt"))
}
