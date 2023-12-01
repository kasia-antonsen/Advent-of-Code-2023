package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
}

func isDigit(character rune) bool {
	return unicode.IsDigit(character)
}
