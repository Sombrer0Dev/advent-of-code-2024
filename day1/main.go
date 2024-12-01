package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Sombrer0Dev/advent-of-code-2024/day1/part1"
	"github.com/Sombrer0Dev/advent-of-code-2024/day1/part2"
)

const day = "day1"

func main() {
	var scanner *bufio.Scanner
	var part int
	var sample bool

	path := os.Getenv("ROOT_PATH")

	fmt.Println(day)

	flag.IntVar(&part, "part", 1, "Specify part of a day")
	flag.BoolVar(&sample, "sample", false, "Use sample input?")
	flag.Parse()

	var file *os.File
	var err error
	if sample {
		file, err = os.Open(path + day + "/sample_" + strconv.Itoa(part))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open(path + day + "/input")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

	switch part {
	case 1:
		fmt.Println("Part1")
		fmt.Printf("Answer: %d \n", part1.Solve(scanner))
	case 2:
		fmt.Println("Part2")
		fmt.Printf("Answer: %d \n", part2.Solve(scanner))
	}
}
