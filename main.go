package main

import (
	"fmt"
)

func coordinate(index int) (r int, c int) {
	// if index == 1{
	// 	return 0,0
	// } else if
	
	switch index {
		case 1:
			fmt.Println(0,0)
			return 0, 0
		case 2:
			fmt.Println(0,0)
			return 0, 1
		case 3:
			fmt.Println(0,0)
			return 0, 2
		case 4:
			fmt.Println(0,0)
			return 1, 0
		case 5:
			fmt.Println(0,0)
			return 1, 1
		case 6:
			fmt.Println(0,0)
			return 1, 2
		case 7:
			fmt.Println(0,0)
			return 2, 0
		case 8:
			fmt.Println(0,0)
			return 2, 1
		case 9:
			fmt.Println(0,0)
			return 2, 2
		default:
			fmt.Println("Invalid input")
			return -1, -1
	}
}

func main() {
	fmt.Println("Hello World!")
	// var board [3][3]int
	position := 0
	fmt.Println("P1 (1-9): ")
	fmt.Scanln(&position)
	x, y := coordinate(position)
	var isDone bool = false
	fmt.Println("p:",position,"x:", x,"y:", y)
	for !isDone {
		p1 := 0
		p2 := 0
		fmt.Println("P1:")
		fmt.Scanln(&p1)
		fmt.Println("P2:")
		fmt.Scanln(&p2)

	}
}