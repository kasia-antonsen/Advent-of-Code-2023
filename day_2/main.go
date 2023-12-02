package main

import (
	util "aoc/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var redRegex = regexp.MustCompile("[1-9][0-9]{0,} red")
var blueRegex = regexp.MustCompile("[1-9][0-9]{0,} blue")
var greenRegex = regexp.MustCompile("[1-9][0-9]{0,} green")

var red = 12
var green = 13
var blue = 14

func main() {
	file, scanner := util.ReadFile("2.txt")
	gamesIdsSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ":")
		title := splitLine[0]
		content := splitLine[1]

		valid := true
		splitContent := strings.Split(content, ";")
		for _, part := range splitContent {
			for _, reach := range strings.Split(part, ",") {
				if !isRedValid(reach) || !isGreenValid(reach) || !isBlueValid(reach) {
					valid = false
				}
			}
		}
		if valid {
			gameNumber, err := strconv.Atoi(strings.Split(title, " ")[1])
			if err == nil {
				fmt.Println(gameNumber)
				gamesIdsSum += gameNumber
			}
		}

	}
	util.CloseFile(file)
	fmt.Println(gamesIdsSum)
}

func isRedValid(part string) bool {
	redString := redRegex.FindString(part)
	if len(redString) == 0 {
		return true
	}
	redValue, err := strconv.Atoi(strings.Split(redString, " ")[0])
	if err != nil || redValue <= red {
		return true
	}

	return false
}

func isBlueValid(part string) bool {
	blueString := blueRegex.FindString(part)
	if len(blueString) == 0 {
		return true
	}
	blueValue, err := strconv.Atoi(strings.Split(blueString, " ")[0])
	if err != nil || blueValue <= blue {
		return true
	}

	return false
}

func isGreenValid(part string) bool {
	greenString := greenRegex.FindString(part)
	if len(greenString) == 0 {
		return true
	}
	greenValue, err := strconv.Atoi(strings.Split(greenString, " ")[0])
	if err != nil || greenValue <= green {
		return true

	}
	return false
}
