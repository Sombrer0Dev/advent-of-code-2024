package part2

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Sombrer0Dev/advent-of-code-2024/utils"
)

var (
	inc   bool
	f_inc bool
	good  int
)

func solve(line []string) bool {
	fmt.Println(line)
	for i := 1; i < len(line); i++ {
		p_lvl, err := strconv.Atoi(line[i-1])
		if err != nil {
			log.Fatal(err)
			return false
		}
		c_lvl, err := strconv.Atoi(line[i])
		if err != nil {
			log.Fatal(err)
			return false
		}

		if i == 1 {
			f_inc = p_lvl > c_lvl
		}
		inc = p_lvl > c_lvl
		distance := utils.Abs(p_lvl - c_lvl)
		if inc != f_inc {
			fmt.Println("inc")
			return false
		}
		if distance < 1 || distance > 3 {
			fmt.Printf("dist error: d=%d, p=%d, c=%d\n", distance, p_lvl, c_lvl)
			return false
		}
	}
	return true
}

func Solve(scanner *bufio.Scanner) int {
	var ans int
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		levels := strings.Split(text, " ")

		var good bool
		for i := -1; i < len(levels); i++ {
			if i == -1 {
				if solve(levels) {
					good = true
					break
				}
			} else {
				if solve(utils.DelSlice(levels, i)) {
					good = true
					break
				}
			}
		}
		if good {
			fmt.Println("GOOD")
		} else {
			fmt.Println("BAD")
		}
		ans += utils.Btoi(good)
	}
	return ans
}
