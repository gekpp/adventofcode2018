package main

import (
	"fmt"
	"os"
)

const (
	mx = 300
	my = 300
)

func main() {

	var serialN int
	_, err := fmt.Fscanf(os.Stdin, "%d", &serialN)
	if err != nil {
		panic(err)
	}

	m := buildMatrix(serialN)
	x, y, p := findMaxPower(m, 3)
	fmt.Printf("Window 3x3: (%d, %d) -> %d", x, y, p)
}

func buildMatrix(serialN int) [][]int {
	res := make([][]int, my, my)
	for y := 1; y <= my; y++ {
		row := make([]int, mx, mx)
		for x := 1; x <= mx; x++ {
			rackId := x + 10
			row[x-1] = ((rackId*y+serialN)*rackId/100)%10 - 5
		}
		res[y-1] = row
	}
	return res
}

func findMaxPower(m [][]int, side int) (maxX int, maxY int, maxP int) {
	for y := 0; y < my-side; y++ {
		for x := 0; x < mx-side; x++ {
			if power := getPower(m, y, x, side); power > maxP {
				maxP = power
				maxX = x
				maxY = y
			}
		}
	}
	return maxX + 1, maxY + 1, maxP
}

func getPower(m [][]int, y, x, side int) (res int) {
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			res += m[y+i][x+j]
		}
	}
	return res
}
