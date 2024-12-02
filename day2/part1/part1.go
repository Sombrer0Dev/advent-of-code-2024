package part1

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Sombrer0Dev/advent-of-code-2024/utils"
)

func Solve(scanner *bufio.Scanner) int {
	var ans int
	for scanner.Scan() {
		var inc bool
		var f_inc bool
		var good int
		// first := true
		text := scanner.Text()
		fmt.Println(text)
		levels := strings.Split(text, " ")
		for i := 1; i < len(levels); i++ {
			good = 0
			p_lvl, err := strconv.Atoi(levels[i-1])
			if err != nil {
				log.Fatal(err)
				return 0
			}
			c_lvl, err := strconv.Atoi(levels[i])
			if err != nil {
				log.Fatal(err)
				return 0
			}
			if i == 1 {
				f_inc = p_lvl > c_lvl
			}
			inc = p_lvl > c_lvl

			if inc != f_inc {
				break
			}

			distance := utils.Abs(p_lvl - c_lvl)
			if distance < 1 || distance > 3 {
				break
			}
			good = 1
		}
		if good == 1 {
			fmt.Println("GOOD")
		} else {
			fmt.Println("BAD")
		}
		ans += good
	}
	return ans
}
