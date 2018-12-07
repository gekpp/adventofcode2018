package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type matrix map[int][]int

type row struct {
	n string
	x int
	y int
	w int
	h int
}

func main() {

	r := bufio.NewReader(os.Stdin)
	plot := make(matrix)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, "\n")

		claim := parseClaim(line)
		updateMatrix(plot, claim.x, claim.y, claim.x+claim.w, claim.y+claim.h)
	}

	var res int
	for _, row := range plot {
		for _, cell := range row {
			if cell > 1 {
				res ++
			}
		}
	}

	fmt.Println(res)
}

func updateMatrix(plot matrix, x0, y0, x1, y1 int) {
	for j := y0; j < y1; j++ {
		row, ok := plot[j]
		if !ok {
			row = make([]int, x1, x1)
		} else if len(row) < x1 {
			delta := x1 - len(row)
			row = append(row, make([]int, delta, delta)...)
		}
		plot[j] = row

		for i := x0; i < x1; i++ {
			row[i] += 1
		}
	}
}

func parseClaim(line string) row {
	prts := strings.Split(line, " ")

	coord := strings.Split(prts[2][:len(prts[2])-1], ",")
	x, _ := strconv.ParseInt(coord[0], 10, 32)
	y, _ := strconv.ParseInt(coord[1], 10, 32)

	dims := strings.Split(prts[3], "x")
	w, _ := strconv.ParseInt(dims[0], 10, 32)
	h, _ := strconv.ParseInt(dims[1], 10, 32)
	return row{prts[0], int(x), int(y), int(w), int(h)}
}
