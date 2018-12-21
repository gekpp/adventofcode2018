package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type plan map[complex64]int

var dirs = map[byte]complex64{'N': -1, 'E': 1i, 'S': 1, 'W': -1i}

func printMap(b int, _ bool) {
	switch b {
	case 0:
		fmt.Print("#")
	default:
		fmt.Printf("%c", b)
	}
}

func main() {
	p := readInput("day20/input-ark.txt")
	printPlan(p, printMap)
	dist := fillDistance(p)

	var max = 0
	var n = 0
	for _, d := range dist {
		if d > max {
			max = d
		}

		if d >= 1000 {
			n++
		}
	}
	fmt.Println(max)
	fmt.Println(n)

}

func fillDistance(p plan) plan {

	res := map[complex64]int{0: 0}
	points := []complex64{0}
	var dist int
	for len(points) > 0 {
		l := len(points)
		for i, pt := range points {
			_ = i
			for _, d := range dirs {
				if p[pt+d] != 0 {
					if _, ok := res[pt+2*d]; !ok {
						res[pt+2*d] = dist + 1
						points = append(points, pt+2*d)
					}
				}
			}
		}
		points = points[l:]
		dist++
	}
	return res
}

func readInput(filename string) plan {

	var (
		pos complex64 = 0
		p             = make(plan)
	)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	pos = readPath(r, pos, p)
	return p
}

func readPath(r *bufio.Reader, pos complex64, p plan) complex64 {
	var (
		checkpoint   = pos
		slashStarted = false
	)

	for {
		//printPlan(p)
		b, err := r.ReadByte()
		if err != nil {
			break
		}

		fmt.Printf("%c", b)
		switch b {
		case '^':
			p[pos] = 0
		case '$':
			break
		case ')':
			if slashStarted {
				return pos
			}
			return checkpoint
		case '|':
			slashStarted = true
		case 'N', 'E', 'S', 'W':
			if slashStarted {
				pos = checkpoint
				slashStarted = false
			}

			p[pos] = ' '
			nextPos := pos + dirs[b]
			if b == 'S' || b == 'N' {
				p[nextPos] = '-'
			} else {
				p[nextPos] = '|'
			}
			pos = nextPos + dirs[b]
			p[pos] = ' '
		case '(':
			pos = readPath(r, pos, p)
		}
	}
	return pos
}

func printPlan(p plan, printFn func(int, bool)) {
	var (
		minX, minY float32 = math.MaxFloat32, math.MaxFloat32
		maxX, maxY float32 = -math.MaxFloat32, -math.MaxFloat32
	)

	for key := range p {
		if real(key) > maxY {
			maxY = real(key)
		}
		if real(key) < minY {
			minY = real(key)
		}

		if imag(key) > maxX {
			maxX = imag(key)
		}
		if imag(key) < minX {
			minX = imag(key)
		}
	}

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if complex(y, x) == 0 {
				fmt.Print("X")
				continue
			}
			b, ok := p[complex(y, x)]
			printFn(b, ok)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
