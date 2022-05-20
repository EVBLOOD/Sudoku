package main

import (
	"fmt"
	"os"
	"strconv"
)

func Split(s string) []rune {
	var result []rune = make([]rune, 0)
	for i := 0; i < len(s); i++ {
		result = append(result, rune(s[i : i+1][0]))
	}
	result = append(result, rune(s[len(s)-2 : len(s)-1][0]))
	return result
}

func checker_one() (bool, [][]rune) {
	var checker [10]rune = [10]rune{'-', '-', '-', '-', '-', '-', '-', '-', '-', '-'}
	var soduko [][]rune = make([][]rune, 0)
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return false, nil
	}
	for i := 1; i < len(os.Args); i++ {
		if len(os.Args[i]) != 9 {
			fmt.Println("Error")
			return false, nil
		}
		for j := 0; j < len(os.Args[i]); j++ {
			val, err := strconv.Atoi(string(os.Args[i][j]))
			val *= 1
			if os.Args[i][j] == '.' || os.Args[i][j] == ' ' || err == nil {
				continue
			} else {

				fmt.Println("Error : " + string(os.Args[i][j]))
				return false, nil
			}
		}
	}
	for i := 1; i < len(os.Args); i++ {
		soduko = append(soduko, Split((os.Args[i])))
		for j := 0; j < 9; j++ {
			if soduko[i-1][j] != '.' {
				p, _ := strconv.Atoi(string(soduko[i-1][j]))
				if checker[p] != rune(p) {
					checker[p] = rune(p)
				} else {
					fmt.Println("Error")
					return false, nil
				}
			}
		}
		checker = [10]rune{'-', '-', '-', '-', '-', '-', '-', '-', '-', '-'}
	}
	return true, soduko
}

func check_x_y(value rune, x int, y int, soduko [][]rune) bool {
	for i := 0; i < 9; i++ {
		if i != y && soduko[x][i] == value {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if i != x && soduko[i][y] == value {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (i != x/3) && (j != y/3) && soduko[i+x/3*3][j+y/3*3] == value {
				return false
			}
		}
	}

	return true
}

func solving(soduko *[][]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*soduko)[i][j] == '.' {
				for k := 1; k < 10; k++ {
					if check_x_y(rune(k+'0'), i, j, *soduko) {
						(*soduko)[i][j] = rune(k + '0')
						if solving(soduko) {
							return true
						}
						(*soduko)[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	check, soduko := checker_one()
	if check == false {
		return
	}
	ok := solving(&soduko)
	if ok {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%c", soduko[i][j])
				if j < 8 {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("error")
	}
}
