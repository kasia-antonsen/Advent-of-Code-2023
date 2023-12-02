package main

import (
	util "aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, scanner := util.ReadFile("1.txt")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		indexOfFirstDigit := strings.IndexFunc(line, isDigit)
		indexOfLastDigit := strings.LastIndexFunc(line, isDigit)

		number := string(line[indexOfFirstDigit]) + string(line[indexOfLastDigit])
		converted, err := strconv.Atoi(number)
		if err == nil {
			sum += converted
		}
	}

	fmt.Print(sum)
	util.CloseFile(file)
}

func isDigit(character rune) bool {
	return unicode.IsDigit(character)
}
