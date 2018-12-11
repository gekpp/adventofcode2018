package main

import (
	"fmt"
	"os"
	"time"
)

const (
	mx = 300
	my = 300
)

func main() {
	fmt.Println(time.Now())
	var serialN int
	_, err := fmt.Fscanf(os.Stdin, "%d", &serialN)
	if err != nil {
		panic(err)
	}

	m := buildMatrix(serialN)

	x, y, p, s := findMaxWindow(m)
	fmt.Printf("Window %dx%d: (%d, %d) -> %d\n", s, s, x, y, p)
	fmt.Println(time.Now())
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

func findMaxWindow(m [][]int) (maxX, maxY, maxP, maxS int) {
	state := make([][]int, my, my)
	for i := 0; i < my; i++ {
		state[i] = make([]int, mx, mx)
		copy(state[i], m[i])
	}
	for s := 2; s <= mx; s++ {
		state = fillPower(state, m, s)
		if x, y, p := findMaxPower(state, s); p > maxP {
			maxX, maxY, maxP, maxS = x, y, p, s
		}
	}
	return maxX, maxY, maxP, maxS
}

func fillPower(prevState, orig [][]int, side int) [][]int {
	res := make([][]int, my, my)
	for y := 0; y < my-side; y++ {
		row := make([]int, mx, mx)
		for x := 0; x < mx-side; x++ {
			row[x] = getPower(prevState, orig, y, x, side)
		}
		res[y] = row
	}
	return res
}

func findMaxPower(state [][]int, side int) (maxX, maxY, maxP int) {
	for y := 0; y < my-side; y++ {
		for x := 0; x < mx-side; x++ {
			if state[y][x] > maxP {
				maxX, maxY, maxP = x+1, y+1, state[y][x]
			}
		}
	}
	return maxX, maxY, maxP
}

func getPower(state, orig [][]int, y, x, side int) (res int) {
	res = state[y][x]
	for i := 0; i < side-1; i++ {
		res += orig[y+i][x+side-1]
	}

	for i := 0; i < side; i++ {
		res += orig[y+side-1][x+i]
	}
	return res
}
