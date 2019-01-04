package main

import (
	"bufio"
	"fmt"
	"hash/crc64"
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

type brick struct {
	points [8]point
}

func (c brick) double(axis int) (res []brick) {
	w := c.width(axis)
	newW := w / 2
	if w > 1 {
		newW += w % 2
	}

	var (
		wx, wy, wz = c.width(0), c.width(1), c.width(2)
	)

	p0 := c.points[0]
	var p1 point
	switch axis {
	case axX:
		wx = newW
		p1 = point{p0.x + newW, p0.y, p0.z}
	case axY:
		wy = newW
		p1 = point{p0.x, p0.y + newW, p0.z}
	case axZ:
		wz = newW
		p1 = point{p0.x, p0.y, p0.z + newW}
	}

	return []brick{
		buildBrick(p0, wx, wy, wz),
		buildBrick(p1, wx, wy, wz),
	}
}

func (c brick) width(axis int) int {
	switch axis {
	case axX:
		return c.points[1].x - c.points[0].x + 1
	case axY:
		return c.points[2].y - c.points[0].y + 1
	case axZ:
		return c.points[4].z - c.points[0].z + 1
	default:
		panic("wrong axis")
	}
}

func buildBrick(min point, wx, wy, wz int) brick {
	return brick{
		[8]point{
			min, {min.x + wx - 1, min.y, min.z}, {min.x + wx - 1, min.y + wy - 1, min.z}, {min.x, min.y + wy - 1, min.z},
			{min.x, min.y, min.z + wz - 1}, {min.x + wx - 1, min.y, min.z + wz - 1}, {min.x + wx - 1, min.y + wy - 1, min.z + wz - 1}, {min.x, min.y + wy - 1, min.z + wz - 1},
		},
	}
}

type bot struct {
	point
	r int
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

func (b bot) interfere1(c brick) bool {

	if inside(b.point, c) {
		return true
	}

	for _, p := range c.points {
		if d := dist(p, b.point); d <= b.r {
			return true
		}
	}

	if botInterfereRectSide(b.point, b.r, c.points[0], c.points[2]) {
		return true
	}

	if botInterfereRectSide(b.point, b.r, c.points[4], c.points[6]) {
		return true
	}

	if botInterfereRectSide(b.point, b.r, c.points[0], c.points[5]) {
		return true
	}

	if botInterfereRectSide(b.point, b.r, c.points[3], c.points[6]) {
		return true
	}

	if botInterfereRectSide(b.point, b.r, c.points[0], c.points[7]) {
		return true
	}

	if botInterfereRectSide(b.point, b.r, c.points[1], c.points[6]) {
		return true
	}

	return false
}

//func (b bot) interfere(c brick) bool {
//
//	if inside(b.point, c) {
//		return true
//	}
//
//	for _, p := range c.points {
//		if d := dist(p, b.point); d <= b.r {
//			return true
//		}
//	}
//
//	if zok := botInterfereRectSide(bot{point{b.x, b.y, 0}, b.r}, project(c.points[0:4], axZ)); !zok {
//		return false
//	}
//
//	if yok := botInterfereRectSide(bot{point{b.x, b.z, 0}, b.r}, project([]point{c.points[0], c.points[1], c.points[5], c.points[4]}, axY)); !yok {
//		return false
//	}
//
//	if xok := botInterfereRectSide(bot{point{b.y, b.z, 0}, b.r}, project([]point{c.points[0], c.points[3], c.points[7], c.points[4]}, axX)); !xok {
//		return false
//	}
//
//	return true
//}

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

	queue := []brick{buildBrick(point{mnx, mny, mnz}, wx, wy, wz)}

	r := 0
	var max int
	for wx > 1 || wy > 1 || wz > 1 {
		b0 := queue[0]
		wx, wy, wz = b0.width(axX), b0.width(axY), b0.width(axZ)
		fmt.Printf("wx,wy,wz=%d, %d, %d. Len=%d.", wx, wy, wz, len(queue))
		longestAxis := maxAxis(wx, wy, wz)
		oldQ := queue
		queue = make([]brick, 0, len(oldQ)*2)
		for _, b := range oldQ {
			queue = append(queue, b.double(longestAxis)...)
		}

		queue, max = searchDown(queue, bots)
		fmt.Printf(" Max=%d\n", max)
		r++
	}

	return queue[0].points[0]
}

func maxAxis(wx, wy, wz int) int {
	switch m := max(wx, max(wy, wz)); m {
	case wx:
		return axX
	case wy:
		return axY
	case wz:
		return axZ
	default:
		panic("bad argument")
	}
}

func searchDown(queue []brick, bots []bot) ([]brick, int) {

	var max int
	cubeToBotsN := make(map[brick]int)
	brickToBotsHash := make(map[brick]uint64)
	var notInterfereBots []int
	for _, c := range queue {
		buf := make([]byte, 0)
		bc := 0
		for i, b := range bots {
			if b.interfere1(c) {
				bc++
				buf = append(buf, []byte(fmt.Sprintf("%04d", i))...)
			} else {
				notInterfereBots = append(notInterfereBots, i)
			}
		}

		cubeToBotsN[c] = bc
		hash := crc64.Checksum(buf, crc64.MakeTable(crc64.ECMA))
		brickToBotsHash[c] = hash
		if bc > max {
			max = bc
		}
	}

	for c, bc := range cubeToBotsN {
		//fmt.Printf("%+v => %d\n", c, bc)
		if bc != max {
			delete(cubeToBotsN, c)
		}
	}
	//fmt.Println()

	queue = make([]brick, 0, len(cubeToBotsN))
	for c := range cubeToBotsN {
		queue = append(queue, c)
	}

	var zero point
	sort.Slice(queue, func(i, j int) bool {
		c0, c1 := queue[i], queue[j]
		return dist(c0.points[0], zero) < dist(c1.points[0], zero)
	})

	hashAdded := make(map[uint64]struct{})
	res := make([]brick, 0)
	for _, c := range queue {
		hash := brickToBotsHash[c]
		if _, ok := hashAdded[hash]; ok {
			continue
		}
		res = append(res, c)
		hashAdded[hash] = struct{}{}
	}
	return res, max
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

func inside(p point, c brick) bool {
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
