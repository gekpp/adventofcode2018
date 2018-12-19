package main

import (
	"bufio"
	"fmt"
	"os"
)

const SIDE = 50

type plot [SIDE][SIDE]byte

func main() {
	original := readInput("day18/input-ark.txt")

	printWorld(original)
	w := original
	for i := 0; i < 10; i++ {
		w = tick(w)
	}
	printWorld(w)
	fmt.Println(count(w, '|') * count(w, '#'))

	w = original
	seen := map[plot]int{w: 0}
	end := 1000000000
	i := 1
	for ; i <= end; i++ {
		w = tick(w)
		if prevStep, ok := seen[w]; ok {
			remain := end - i
			period := i - prevStep
			nextI := i + (remain/period)*period
			for i:= nextI+1; i <= end; i++ {
				w = tick(w)
			}
			break
		}
		seen[w] = i
	}

	printWorld(w)
	fmt.Println(count(w, '|') * count(w, '#'))
}

func tick(p plot) (res plot) {
	for y, row := range p {
		for x, c := range row {
			switch c {
			case '.':
				res[y][x] = handleOpen(p, y, x)
			case '|':
				res[y][x] = handleTrees(p, y, x)
			case '#':
				res[y][x] = handleLumberyard(p, y, x)
			}
		}
	}
	return res
}

func count(p plot, des byte) (res int) {
	for _, row := range p {
		for _, c := range row {
			if c == des {
				res ++
			}
		}
	}
	return res
}

func handleOpen(p plot, y, x int) byte {
	if countDesired(p, y, x, '|') >= 3 {
		return '|'
	}
	return '.'
}

func handleTrees(p plot, y, x int) byte {
	if countDesired(p, y, x, '#') >= 3 {
		return '#'
	}
	return '|'
}

func handleLumberyard(p plot, y, x int) byte {
	if countDesired(p, y, x, '#') >= 1 && countDesired(p, y, x, '|') >= 1 {
		return '#'
	}
	return '.'
}

func countDesired(p plot, y, x int, desire byte) int {
	count := 0
	for y1 := y - 1; y1 <= y+1; y1++ {
		if y1 < 0 || y1 >= SIDE {
			continue
		}
		for x1 := x - 1; x1 <= x+1; x1++ {
			if x1 < 0 || x1 >= SIDE || (y1 == y && x1 == x) {
				continue
			}
			if p[y1][x1] == desire {
				count++
			}
		}
	}
	return count
}

func printWorld(w [SIDE][SIDE]byte) {
	for _, row := range w {
		fmt.Printf("%s\n", row)
	}
	fmt.Println()
	fmt.Println()
}

func readInput(filename string) (res [SIDE][SIDE]byte) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	ln := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		bytes := s.Bytes()
		for i, b := range bytes {
			res[ln][i] = b
		}
		ln++
	}

	return res
}
