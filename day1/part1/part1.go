package part1

import (
	"bufio"
	"fmt"
)

func Solve(scanner *bufio.Scanner) {
	fmt.Println("Part1")
	for scanner.Scan(){
		text := scanner.Text()
		fmt.Println(text)
	}
}
