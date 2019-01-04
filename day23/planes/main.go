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

const (
	axX = iota
	axY
	axZ
)

type point struct {
	x, y, z int
}

type bot struct {
	point
	r int
}

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func main() {
	//bots := readInput("day23/input-test-2.txt")
	bots := readInput("day23/input-ark.txt")
	usi := underStrongestInfluence(bots)
	fmt.Printf("There are %d nanobots in range\n", usi)

	min := findMax(bots)
	fmt.Println(min)
	fmt.Println(dist(min, point{}))
}

func findMax(bots []bot) point {
	sort.Slice(bots, func(i, j int) bool {
		return bots[i].x < bots[j].x
	})
	mnx := bots[0].x
	mxx := bots[len(bots)-1].x

	sort.Slice(bots, func(i, j int) bool {
		return bots[i].y < bots[j].y
	})
	mny := bots[0].y
	mxy := bots[len(bots)-1].y

	sort.Slice(bots, func(i, j int) bool {
		return bots[i].z < bots[j].z
	})
	mnz := bots[0].z
	mxz := bots[len(bots)-1].z

	fmt.Println("\nOptimizing z...")
	z := searchDown(bots, axZ, mnx, mxx, mny, mxy, mnz, mxz)
	fmt.Println("\nOptimizing x...")

	x := searchDown(bots, axX, mnx, mxx, mny, mxy, mnz, mxz)
	fmt.Println("\nOptimizing y...")
	y := searchDown(bots, axY, x, x, mny, mxy, mnz, mxz)

	return point{x, y, z}
}

func sortFn(axis int, bots []bot) func(i, j int) bool {
	switch axis {
	case 0:
		return func(i, j int) bool {
			bi := bots[i]
			bj := bots[j]
			return bi.x-bi.r < bj.x-bj.r
		}
	case 1:
		return func(i, j int) bool {
			bi := bots[i]
			bj := bots[j]
			return bi.y-bi.r < bj.y-bj.r
		}
	case 2:
		return func(i, j int) bool {
			bi := bots[i]
			bj := bots[j]
			return bi.z-bi.z < bj.z-bj.r
		}
	default:
		panic("bad axis")
	}
}

func tooFar(axis int, b bot, coord int) bool {
	switch axis {
	case axX:
		return b.x-b.r > coord
	case axY:
		return b.y-b.r > coord
	case axZ:
		return b.z-b.r > coord
	default:
		panic("bad axis")
	}
}

func searchDown(bots []bot, axis int, mnx, mxx, mny, mxy, mnz, mxz int) int {

	var min, max int
	switch axis {
	case axX:
		min, max = mnx, mxx
	case axY:
		min, max = mny, mxy
	case axZ:
		min, max = mnz, mxz
	}

	sort.Slice(bots, sortFn(axis, bots))
	maxRes := 0
	results := make(map[int]int)
	for coord := min; coord <= max; coord += 3 {
		p0, p2 := buildPlane(axis, coord, mnx, mxx, mny, mxy, mnz, mxz)

		var bc int
		for _, b := range bots {
			if tooFar(axis, b, coord) {
				break
			}
			fmt.Println(b)
			if botInterfereRectSide(b.point, b.r, p0, p2) {
				bc++
			}
		}
		if bc < maxRes || bc == 0 {
			continue
		}
		if bc > maxRes {
			maxRes = bc
		}
		results[coord] = bc
	}

	var res []int
	for coord, count := range results {
		if count == maxRes {
			res = append(res, coord)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res[0]
}

func buildPlane(axis, coord, mnx, mxx, mny, mxy, mnz, mxz int) (p0, p2 point) {
	switch axis {
	case axX:
		return point{coord, mny, mnz}, point{coord, mxy, mxz}
	case axY:
		return point{mnx, coord, mnz}, point{mxx, coord, mxz}
	case axZ:
		return point{mnx, mny, coord}, point{mxx, mxy, coord}
	default:
		panic("bad axis")
	}
}

func botInterfereRectSide(p point, r int, p0, p2 point) bool {
	var nearest point

	if between(p.x, p0.x, p2.x) {
		nearest.x = p.x
	} else {
		if abs(p.x-p0.x) < abs(p.x-p2.x) {
			nearest.x = p0.x
		} else {
			nearest.x = p2.x
		}
	}

	if between(p.y, p0.y, p2.y) {
		nearest.y = p.y
	} else {
		if abs(p.y-p0.y) < abs(p.y-p2.y) {
			nearest.y = p0.y
		} else {
			nearest.y = p2.y
		}
	}

	if between(p.z, p0.z, p2.z) {
		nearest.z = p.z
	} else {
		if abs(p.z-p0.z) < abs(p.z-p2.z) {
			nearest.z = p0.z
		} else {
			nearest.z = p2.z
		}
	}

	return dist(nearest, p) <= r
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

func between(v, b1, b2 int) bool {
	if b1 <= b2 {
		return b1 <= v && v <= b2
	} else {
		return b2 <= v && v <= b1
	}
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
