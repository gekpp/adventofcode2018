package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type point struct {
	x, y, z int
}

const (
	axX = iota
	axY
	axZ
)

type cube struct {
	points [8]point
}

func (c cube) split() (res []cube) {
	width := c.width() / 2
	if c.width() > 1 {
		width += c.width() % 2
	}
	p0 := c.points[0]
	return []cube{
		buildCube(p0, width),
		buildCube(point{p0.x + width, p0.y, p0.z}, width),
		buildCube(point{p0.x + width, p0.y + width, p0.z}, width),
		buildCube(point{p0.x, p0.y + width, p0.z}, width),

		buildCube(point{p0.x, p0.y, p0.z + width}, width),
		buildCube(point{p0.x + width, p0.y, p0.z + width}, width),
		buildCube(point{p0.x + width, p0.y + width, p0.z + width}, width),
		buildCube(point{p0.x, p0.y + width, p0.z + width}, width),
	}
}

func (c cube) width() int {
	return c.points[1].x - c.points[0].x + 1
}

func buildCube(min point, side int) cube {
	return cube{
		[8]point{
			min, {min.x + side - 1, min.y, min.z}, {min.x + side - 1, min.y + side - 1, min.z}, {min.x, min.y + side - 1, min.z},
			{min.x, min.y, min.z + side - 1}, {min.x + side - 1, min.y, min.z + side - 1}, {min.x + side - 1, min.y + side - 1, min.z + side - 1}, {min.x, min.y + side - 1, min.z + side - 1},
		},
	}
}

type bot struct {
	point
	r int
}

//
//func (b bot) interfere(c cube) bool {
//	if inside(b.point, c) {
//		return true
//	}
//
//	for _, p := range c.points {
//		if dist(p, b.point) <= b.r {
//			return true
//		}
//	}
//	return false
//}

func (b bot) interfere(c cube) bool {

	if inside(b.point, c) {
		return true
	}

	for _, p := range c.points {
		if d := dist(p, b.point); d <= b.r {
			return true
		}
	}

	if zok := botInterfereRectSide(bot{point{b.x, b.y, 0}, b.r}, project(c.points[0:4], axZ)); !zok {
		return false
	}

	if yok := botInterfereRectSide(bot{point{b.x, b.z, 0}, b.r}, project([]point{c.points[0], c.points[1], c.points[5], c.points[4]}, axY)); !yok {
		return false
	}

	if xok := botInterfereRectSide(bot{point{b.y, b.z, 0}, b.r}, project([]point{c.points[0], c.points[3], c.points[7], c.points[4]}, axX)); !xok {
		return false
	}

	return true
}

func botInterfereRectSide(b bot, p []point) bool {

	if between(b.x, p[0].x, p[1].x) {
		return between(b.y, p[0].y, p[3].y) || abs(b.y-p[0].y) <= b.r || abs(b.y-p[2].y) <= b.r
	}

	if between(b.y, p[0].y, p[3].y) {
		return abs(b.x-p[0].x) <= b.r || abs(b.x-p[1].x) <= b.r
	}
	return false
}

func between(v, b1, b2 int) bool {
	if b1 <= b2 {
		return b1 <= v && v <= b2
	} else {
		return b2 <= v && v <= b1
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func project(points []point, zeroAxis int) []point {
	res := make([]point, len(points), len(points))
	for i, p := range points {
		switch zeroAxis {
		case 2:
			res[i] = point{x: p.x, y: p.y}
		case 1:
			res[i] = point{x: p.x, y: p.z}
		case 0:
			res[i] = point{x: p.y, y: p.z}
		}
	}
	return res
}

func main() {
	//bots := readInput("day23/input-test-2.txt")
	bots := readInput("day23/input-ark.txt")
	usi := underStrongestInfluence(bots)
	fmt.Printf("There are %d nanobots in range\n", usi)

	min := findMin(bots)
	fmt.Println(min)
	fmt.Println(dist(min, point{}))
}

func findMin(bots []bot) point {
	sort.Slice(bots, func(i, j int) bool {
		return bots[i].x < bots[j].x
	})
	mnx := bots[0].x
	mxx := bots[len(bots)-1].x
	wx := mxx - mnx + 1

	sort.Slice(bots, func(i, j int) bool {
		return bots[i].y < bots[j].y
	})
	mny := bots[0].y
	mxy := bots[len(bots)-1].y
	wy := mxy - mny + 1

	sort.Slice(bots, func(i, j int) bool {
		return bots[i].z < bots[j].z
	})
	mnz := bots[0].z
	mxz := bots[len(bots)-1].z
	wz := mxz - mnz + 1

	width := max(max(wx, wy), wz)
	area := buildCube(point{mnx, mny, mnz}, width)

	var zero point
	for width >= 0 {
		var max int
		queue := area.split()
		fmt.Printf("Width=%d\n", queue[0].width())
		botsCount := make(map[cube]int)
		for _, c := range queue {
			bc := 0
			for _, b := range bots {
				if b.interfere(c) {
					bc++
				}
			}

			fmt.Printf("%+v => %d\n", c, bc)
			botsCount[c] = bc
			if bc > max {
				max = bc
			}
		}
		fmt.Printf("\n\n")

		sort.Slice(queue, func(i, j int) bool {
			ci := queue[i]
			cj := queue[j]
			if botsCount[ci] > botsCount[cj] {
				return true
			}
			if botsCount[ci] == botsCount[cj] {
				return dist(ci.points[0], zero) < dist(cj.points[0], zero)
			}
			return false
		})

		area = queue[0]

		if width == 0 {
			break
		}
		width = width / 2
	}
	return area.points[0]
}

func underStrongestInfluence(bots []bot) int {
	sort.Slice(bots, func(i, j int) bool {
		return bots[i].r < bots[j].r
	})
	fmt.Printf("The strongest bot is %v\n", bots[len(bots)-1])

	maxBot := bots[len(bots)-1]
	maxD := maxBot.r
	res := 0
	for i := 0; i < len(bots); i++ {
		if dist(bots[i].point, maxBot.point) <= maxD {
			res++
		}
	}
	return res
}

func dist(b1, b2 point) int {
	return int(math.Abs(float64(b1.x-b2.x)) + math.Abs(float64(b1.y-b2.y)) + math.Abs(float64(b1.z-b2.z)))
}

func inside(p point, c cube) bool {
	var (
		lx, ly, lz int
	)
	for _, cp := range c.points {
		if cp.x <= p.x {
			lx++
		}
		if cp.y <= p.y {
			ly++
		}
		if cp.z <= p.z {
			lz++
		}
	}
	return lx == 4 && ly == 4 && lz == 4
}

func readInput(filename string) (res []bot) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("could not open file:%v", err)
	}

	s := bufio.NewScanner(file)

	for s.Scan() {
		var b bot
		parts := strings.Split(strings.TrimLeft(s.Text(), "pos=<"), ">, r=")
		parts[0] = strings.Replace(parts[0], ",", " ", -1)
		_, err := fmt.Sscanf(parts[0], "%v %v %d", &b.x, &b.y, &b.z)
		if err != nil {
			log.Fatalf("could not scan line: %v", err)
		}

		if _, err := fmt.Sscanf(parts[1], "%v", &b.r); err != nil {
			log.Fatalf("could not scan line: %v", err)
		}
		res = append(res, b)
	}

	return res
}
