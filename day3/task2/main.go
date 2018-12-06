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

	var stdin io.Reader = os.Stdin
	//	var stdin io.Reader = bytes.NewBuffer([]byte(`#1 @ 1,3: 4x4
	//#2 @ 3,1: 4x4
	//#3 @ 5,5: 2x2
	//`))
	r := bufio.NewReader(stdin)
	plot := make(matrix)
	var claims []row

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, "\n")

		claim := parseClaim(line)
		claims = append(claims, claim)
		updateMatrix(plot, claim)
	}

	for _, claim := range claims {
		if ! hasOverlapping(plot, claim){
			fmt.Println(claim.n)
		}
	}
}

func updateMatrix(plot matrix, claim row) {
	x0, y0, x1, y1 := claim.x, claim.y, claim.x+claim.w, claim.y+claim.h
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

func hasOverlapping(plot matrix, claim row) bool {
	x0, y0, x1, y1 := claim.x, claim.y, claim.x+claim.w, claim.y+claim.h
	for j := y0; j < y1; j++ {
		row, _ := plot[j]

		for i := x0; i < x1; i++ {
			if row[i] > 1 {
				return true
			}
		}
	}
	return false
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
