package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func main() {
	points := readInput()
	bound, maxCntSec := moveToMinBound(points)
	_, _ = fmt.Fprintf(os.Stderr, "Elapsed %d seconds\n", maxCntSec)
	draw(points, maxCntSec, bound)
}

func draw(points []point, sec int, bound point) {
	img := image.NewGray(image.Rect(bound.x-10, bound.y-10, bound.vx+bound.x+10, bound.vy+bound.y+10))

	for _, p := range points {
		img.Set(p.x+p.vx*sec, p.y+p.vy*sec, color.Gray{Y: 255})
	}

	if err := png.Encode(os.Stdout, img); err != nil {
		panic(err)
	}
}

func moveToMinBound(points []point) (point, int) {
	p := points[0]
	var (
		maxCnt    int
		maxCntSec int
		minBound  point
	)

	for sec := 0; sec < 100000; sec ++ {
		var (
			minx, miny, maxx, maxy               = p.x+p.vx*sec, p.y+p.vy*sec, p.x+p.vx*sec, p.y+p.vy*sec
			leftCnt, rightCnt, topCnt, bottomCnt = 0, 0, 0, 0
		)

		for _, p := range points {
			//p is a copy of original point
			p.x = p.x + p.vx*sec
			p.y = p.y + p.vy*sec
			if p.x < minx {
				minx = p.x
				leftCnt = 1
			} else if p.x == minx {
				leftCnt++
			}

			if p.x > maxx {
				maxx = p.x
				rightCnt = 1
			} else if p.x == maxx {
				rightCnt++
			}

			if p.y < miny {
				miny = p.y
				topCnt = 1
			} else if p.y == miny {
				topCnt++
			}
			if p.y > maxy {
				maxy = p.y
				bottomCnt = 1
			} else if p.y == maxy {
				bottomCnt++
			}
		}

		boundCnt := leftCnt + rightCnt + topCnt + bottomCnt
		if boundCnt > maxCnt {
			maxCnt = boundCnt
			maxCntSec = sec
			minBound = point{minx, miny, maxx - minx, maxy - miny}
		}
	}

	return minBound, maxCntSec
}

func readInput() []point {
	var (
		stdin io.Reader = os.Stdin
		res   []point
	)

	r := bufio.NewReader(stdin)
	for {
		line, _, err := r.ReadLine()

		if err != nil && err != io.EOF {
			panic(err)
		}

		if err == io.EOF {
			break
		}

		res = append(res, parseLine(string(line)))
	}
	return res
}

func parseLine(line string) point {
	line = strings.Replace(line, " ", "", -1)
	var (
		p point
	)
	_, err := fmt.Sscanf(line, "position=<%v,%v>velocity=<%v,%v>", &p.x, &p.y, &p.vx, &p.vy)
	if err != nil {
		panic(err)
	}

	return p
}
