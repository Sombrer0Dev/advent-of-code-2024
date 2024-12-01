package part1

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/Sombrer0Dev/advent-of-code-2024/utils"
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

	slices.SortFunc(loc_1, func(a, b int) int {
		return a - b
	})
	slices.SortFunc(loc_2, func(a, b int) int {
		return a - b
	})

	for i := range loc_1 {
		num := utils.Abs(loc_1[i] - loc_2[i])

		fmt.Println(num)
		ans += num
	}
	return ans
}
