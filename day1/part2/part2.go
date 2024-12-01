package part2

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func Solve(scanner *bufio.Scanner) int {
	var loc_1, loc_2 []int
	var ans int

	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, "   ")

		l_1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
			return 0
		}
		l_2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
			return 0
		}

		loc_1 = append(loc_1, l_1)
		loc_2 = append(loc_2, l_2)
	}
	for _, v_1 := range loc_1 {
		count := 0
		for _, v_2 := range loc_2 {
			if v_1 == v_2 {
				count++
			}
		}
		ans += count * v_1
	}
	return ans
}
