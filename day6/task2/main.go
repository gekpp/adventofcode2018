package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

const threshold = 10000

func main() {
	var stdin io.Reader
	stdin = os.Stdin
	r := bufio.NewReader(stdin)

	points := readPoints(r)

	field := fillField(points)
	var res int
	for _, d := range field {
		if d < threshold {
			res += 1
		}
	}
	fmt.Printf("Res = %v\n", res)
}

func readPoints(r *bufio.Reader) []point {
	var points []point
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		split := strings.Split(string(line), ", ")
		x, _ := strconv.ParseInt(split[0], 10, 32)
		y, _ := strconv.ParseInt(split[1], 10, 32)
		points = append(points, point{int(x), int(y)})
	}
	return points
}

func fillField(points []point) map[point]int {

	field := make(map[point]int)
	for _, p := range points {
		fillFrom(field, p.x, p.y, points)
	}
	return field
}

func fillFrom(field map[point]int, x, y int, points []point) {
	p0 := point{x, y}
	if _, ok := field[point{x, y}]; ok {
		return
	}

	var d int
	for _, p := range points {
		d += dist(p0, p)
	}
	field[p0] = d

	if d >= threshold {
		return
	}
	fillFrom(field, x-1, y, points)
	fillFrom(field, x, y-1, points)
	fillFrom(field, x+1, y, points)
	fillFrom(field, x, y+1, points)

}

func dist(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}
