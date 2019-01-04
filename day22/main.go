package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	torch = 1 << iota
	climb
	neith
)

var typeToTool = map[int]int{0: 3, 1: 6, 2: 5}

type point struct {
	x, y int
}

func (p point) left() point {
	return point{p.x - 1, p.y}
}

func (p point) right() point {
	return point{p.x + 1, p.y}
}

func (p point) up() point {
	return point{p.x, p.y - 1}
}

func (p point) down() point {
	return point{p.x, p.y + 1}
}

//map tool -> distance
type cell map[int]int
type pathes map[point]cell

func main() {
	depth, target := readInput("day22/input-ark.txt")
	errIndex, geoIndex, max := generateCave(depth, target)
	_ = geoIndex
	types := make(map[point]int)
	updateTypes(types, errIndex, max)
	rl2 := countRiskLevel2(types, target)
	fmt.Println(rl2)

	elapsed := fastestTime(types, errIndex, geoIndex, target, max, depth)
	fmt.Println(elapsed)

}

func fastestTime(types, errIndex, geoIndex map[point]int, target, max point, depth int) int {
	targetRes := getNaiveResult(types, target)

	resMap := make(pathes)
	resMap[point{0, 0}] = cell{torch: 0, climb: 7}

	front := []point{{0, 0}}
	for len(front) > 0 {
		prevLen := len(front)
		for _, p := range front {
			if exceedsMaximum(resMap[p], targetRes) {
				continue
			}

			for _, p1 := range []point{p.down(), p.right(), p.left(), p.up()} {
				if p1.x < 0 || p1.y < 0 {
					continue
				}

				if p.x > max.x {
					addCol(errIndex, geoIndex, target, depth, &max)
					updateTypes(types, errIndex, max)
				}

				if p.y > max.y {
					addRow(errIndex, geoIndex, target, depth, &max)
					updateTypes(types, errIndex, max)
				}

				if newState, better := worthMove(types[p1], resMap[p], resMap[p1], targetRes); better {
					front = append(front, p1)
					resMap[p1] = newState

					if p1 == target {
						targetRes = newState[torch]
					}
				}
			}

		}
		front = front[prevLen:]
	}

	return targetRes
}

func worthMove(t int, from, to cell, maxResult int) (state cell, better bool) {
	var (
		toolFound      int
		commonToolCost int
	)
	if to == nil {
		to = make(cell)
	}

	// find first common tool
	for tool, fromCost := range from {
		if can(t, tool) {
			toolFound = tool
			if prevCost, ok := to[tool]; !ok || prevCost > fromCost+1 {
				commonToolCost = fromCost + 1
				to[tool] = commonToolCost
				better = true
			} else {
				commonToolCost = prevCost
			}
			break
		}
	}

	otherTool := otherTool(t, toolFound)
	var otherToolCost int
	if fromCost, ok := from[otherTool]; ok {
		otherToolCost = fromCost + 1
	} else {
		otherToolCost = commonToolCost + 7
	}

	if prevCost, ok := to[otherTool]; !ok || prevCost > otherToolCost {
		to[otherTool] = otherToolCost
		better = true
	}

	if exceedsMaximum(to, maxResult) {
		return to, false
	}
	return to, better
}

func exceedsMaximum(cell cell, max int) bool {
	for _, cost := range cell {
		if cost < max {
			return false
		}
	}
	return true
}

func getNaiveResult(types map[point]int, target point) int {
	res := 0
	tool := torch
	var cur point
	for cur.x < target.x {
		if can(types[cur.right()], tool) {
			res += 1
		} else {
			tool = otherTool(types[cur], tool)
			res += 8
		}
		cur = cur.right()
	}

	for cur.y < target.y {
		if can(types[cur.down()], tool) {
			res += 1
		} else {
			tool = otherTool(types[cur], tool)
			res += 8
		}
		cur = cur.down()
	}

	if tool != torch {
		res += 7
	}
	return res
}

func otherTool(typ, tool int) int {
	return typeToTool[typ] ^ tool
}

func can(typ, tool int) bool {
	return typeToTool[typ]&tool > 0
}

func countRiskLevel2(types map[point]int, target point) (res int) {
	for y := 0; y <= target.y; y++ {
		for x := 0; x <= target.x; x++ {
			res += types[point{x, y}]
		}
	}
	return res
}

func updateTypes(types, errIndex map[point]int, max point) {
	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			types[point{x, y}] = errIndex[point{x, y}] % 3
		}
	}
}

func readInput(filename string) (depth int, target point) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Scan()
	_, _ = fmt.Sscanf(strings.TrimLeft(s.Text(), "depth: "), "%v", &depth)
	s.Scan()
	_, _ = fmt.Sscanf(strings.TrimLeft(s.Text(), "target: "), "%v,%v", &target.x, &target.y)
	return depth, target
}

func generateCave(depth int, target point) (errIndex, geoIndex map[point]int, max point) {
	geoIndex = make(map[point]int)
	errIndex = make(map[point]int)

	max = point{0, 0}
	for x := 0; x <= target.x; x++ {
		addCol(errIndex, geoIndex, target, depth, &max)
	}

	for y := 0; y <= target.y; y++ {
		addRow(errIndex, geoIndex, target, depth, &max)
	}

	return errIndex, geoIndex, max
}

func addRow(errIndex, geoIndex map[point]int, target point, depth int, max *point) {
	for x := 0; x <= max.x; x++ {
		addPoint(point{x, max.y + 1}, target, depth, errIndex, geoIndex)
	}
	max.y = max.y + 1
}

func addCol(errIndex, geoIndex map[point]int, target point, depth int, max *point) {
	for y := 0; y <= max.y; y++ {
		addPoint(point{max.x + 1, y}, target, depth, errIndex, geoIndex)
	}
	max.x = max.x + 1
}

func addPoint(p, target point, depth int, errIndex, geoIndex map[point]int) {
	var gi int
	switch {
	case p.y == target.y && p.x == target.x:
		gi = 0
	case p.y == 0:
		gi = p.x * 16807
	case p.x == 0:
		gi = p.y * 48271
	default:
		gi = errIndex[point{p.x, p.y - 1}] * errIndex[point{p.x - 1, p.y}]
	}
	geoIndex[p] = gi
	errIndex[p] = (gi + depth) % 20183
}
