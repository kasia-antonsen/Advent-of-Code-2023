package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filename string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file, bufio.NewScanner(file)
}

func CloseFile(file *os.File) {
	file.Close()
}
