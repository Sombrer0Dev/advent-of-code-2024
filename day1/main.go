package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Sombrer0Dev/advent-of-code-2024/day1/part1"
	"github.com/Sombrer0Dev/advent-of-code-2024/day1/part2"
	"github.com/Sombrer0Dev/advent-of-code-2024/utils"
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

	if sample {
		scanner = utils.ReadInput(path + day + "/sample_" + strconv.Itoa(part))
	} else {
		scanner = utils.ReadInput(path + day + "/input")
	}

	switch part {
	case 1:
		part1.Solve(scanner)
	case 2:
		part2.Solve(scanner)
	}

}
