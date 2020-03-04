package main

import (
	"fmt"
	"strconv"
)

func testNumber1A (n int) int {
	if n == 1 { //base case
		return 1
	}

	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")

		if i == 1 {
			fmt.Print("\n")
		}
	}
	return testNumber1A(n-1)
}

func tesNumber1B (n int) {
	//init diagonal for cross #
	diagonalX := 1
	diagonalY := n - 1

	//count how many repeat
	repeat := 1

	for i := 1; i <= n - 1; i++ {
		if repeat == n { //break the loop when it's reach n
			break
		}

		if repeat == 1 || repeat == n - 1 { //if the repeatition is number 1 or the last
			fmt.Print("#")
			if i == n - 1 {
				fmt.Print("\n")
			}
		} else {
			if i == 1 || i == diagonalX || i == diagonalY || i == n - 1 {
				fmt.Print("#")
				if i == n - 1 {
					fmt.Print("\n")
				}
			} else {
				fmt.Print(" ")
			}
		}

		if i == n - 1 { //repeat the process again
			i = 0
			repeat += 1
			diagonalX += 1
			diagonalY -= 1
		}
	}
}

func tesNumber2 (n int) {

	//make a dynamic zero string
	init := strconv.Itoa(n)
	zero := ""
	for x := 0; x < len(init) - 1; x++ { //len - 1 because the x start from 0 (incursion)
		zero += "0"
	}

	//make index for zero string
	index := 0

	for _, item := range init {
		if item == 48 { //skip when item is 0 / 48 byte ascii
			index += 1
			continue
		}
		fmt.Println(string(item) + zero[index:])
		index += 1
	}
}

func main() {
	var n, x int
	fmt.Scan(&n, &x) //scan input number, n for tes number 1. x for tes number 2

	if n <= 0 {
		fmt.Print("must be a positive number or more than 0!")
	} else {
		fmt.Println(testNumber1A(n), "\n")
		tesNumber1B(n)
	}
	fmt.Print("\n")
	tesNumber2(x)
}