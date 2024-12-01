package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return scanner
}
