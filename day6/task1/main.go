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

type marker struct {
	p *point
	d int
}

func main() {
	var stdin io.Reader
	stdin = os.Stdin
	r := bufio.NewReader(stdin)

	points := readPoints(r)
	min, max := getBounds(points)
	field := make([][]*marker, max.y-min.y+1)
	for i := 0; i < max.y-min.y+1; i++ {
		field[i] = make([]*marker, max.x-min.x+1)
	}

	for _, p := range points {
		fillField(field, p)
	}

	skipPoints := getSkipPoints(field)
	var (
		maxD int
		maxP point
	)

	res := make(map[point]int)
	for _, row := range field {
		for _, col := range row {
			if col.p == nil {
				continue
			}

			if _, ok := skipPoints[*col.p]; ok {
				continue
			}

			if col.p != nil {
				res[*col.p] = res[*col.p] + 1
			}
		}
	}

	for p, cnt := range res {
		if cnt > maxD {
			maxD = cnt
			maxP = p
		}
	}

	fmt.Printf("MaxD=%v, MaxP=%+v\n", maxD, maxP)
}

func getSkipPoints(points [][]*marker) map[point]struct{} {
	res := make(map[point]struct{})
	for i := 0; i < len(points[0]); i++ {
		if points[0][i].p != nil {
			res[*points[0][i].p] = struct{}{}
		}
		if points[len(points)-1][i].p != nil {
			res[*points[len(points)-1][i].p] = struct{}{}
		}
	}

	for j := 0; j < len(points); j++ {
		if points[j][0].p != nil {
			res[*points[j][0].p] = struct{}{}
		}

		if points[j][len(points[j])-1].p != nil {
			res[*points[j][len(points[j])-1].p] = struct{}{}
		}
	}
	return res
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

func fillField(field [][]*marker, p point) {
	for i, row := range field {
		for j, col := range row {
			d := dist(p, point{j, i})
			if col == nil {
				field[i][j] = &marker{&p, d}
				continue
			}

			if col.d > d {
				col.p = &p
				col.d = d
				continue
			}

			if col.d == d {
				col.p = nil
			}
		}
	}
}

func dist(p1, p2 point) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func getBounds(points []point) (min, max point) {
	var (
		minX, minY = math.MaxInt32, math.MaxInt32
		maxX, maxY int
	)

	for _, p1 := range points {
		if p1.x < minX {
			minX = p1.x
		}
		if p1.x > maxX {
			maxX = p1.x
		}
		if p1.y < minY {
			minY = p1.y
		}
		if p1.y > maxY {
			maxY = p1.y
		}
	}

	return point{minX, minY}, point{maxX, maxY}
}
