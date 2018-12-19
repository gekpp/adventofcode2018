package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"
	"strings"
)

type world map[complex64]int8

var min, max complex64

func main() {

	f, err := os.Open("day17/input-ark.txt")
	if err != nil {
		panic(err)
	}
	w, min1, max1 := readInput(f)
	min, max = min1, max1

	var (
		miny float32 = math.MaxFloat32
	)
	for p := range w {
		if real(p) < miny {
			miny = real(p)
		}
	}

	_ = min
	drawWorld(w, min, max, "day17/render/1.png")

	step := 1
	fill(w, complex(0, 500), int(real(max)), &step)
	drawWorld(w, min, max, "day17/render/2.png")

	fmt.Println(count(w, miny))

	drains := getMaxPoints(w)
	for _, p := range drains {
		retain(w, p)
	}
	drawWorld(w, min, max, "day17/render/3.png")
	fmt.Println(count(w, miny))
}

func getMaxPoints(w world) []complex64 {
	res := make([]complex64, 0)
	var maxY float32 = 0
	for p, c := range w {
		if c == 1 {
			if real(p) > maxY {
				maxY = real(p)
				res = []complex64{p}
			}
			if real(p) == maxY {
				res = append(res, p)
			}
		}
	}
	return res
}

func retain(w world, p complex64) {
	if w[p] != 1 {
		return
	}

	w[p] = 0
	retain(w, p+1i)
	retain(w, p-1i)
	retain(w, p-1)
}

func count(w world, miny float32) int {
	res := 0
	for p, b := range w {
		if real(p) < miny {
			continue
		}
		if b == 1 {
			res++
		}
	}
	return res
}

func fill(w world, start complex64, maxY int, step *int) bool {
	if w[start] != 0 {
		return false
	}

	if outOfBound(start, maxY) {
		return true
	}

	if w[start] == 0 {
		w[start] = 1
	}

	*step += 1
	out := fill(w, start+1, maxY, step)

	if out {
		return out
	}
	//drawWorld(w, min, max, fmt.Sprintf("day17/render/%03d.png", *step))

	if shallSpread(w, start, -1i) && shallSpread(w, start, 1i) {
		*step += 1
		outL := fill(w, start-1i, maxY, step)
		out = out || outL
		//}
		//drawWorld(w, min, max, fmt.Sprintf("day17/render/%03d.png", *step))

		//if shallSpread(w, start, 1i) {
		*step += 1
		outR := fill(w, start+1i, maxY, step)
		out = out || outR
	}
	//drawWorld(w, min, max, fmt.Sprintf("day17/render/%03d.png", *step))

	//if *step%5 == 0 {
	//drawWorld(w, min, max, fmt.Sprintf("day17/render/%03d.png", *step))
	//}

	return out
}

func shallSpread(w world, p, inc complex64) bool {
	var steps int
	for {
		if imag(p) <= imag(min) || imag(p) >= imag(max) {
			return true
		}

		switch {
		case w[p] == -1 || w[p+1] == -1:
			return true

		}

		if w[p+1] == 0 {
			return false
		}
		p = p + inc
		steps ++
	}
}

func outOfBound(point complex64, maxY int) bool {
	return int(real(point)) > maxY
}

func readInput(r io.Reader) (world, complex64, complex64) {
	w := make(world)
	var (
		min, max complex64
	)

	min = complex(0, 500)
	max = complex(0, 500)

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		var (
			axC       byte
			c, v1, v2 float32
		)

		split := strings.Split(line, ", ")
		axC = split[0][0]
		if _, err := fmt.Sscanf(split[0][2:], "%f", &c); err != nil {
			panic(err)
		}

		split = strings.Split(split[1][2:], "..")

		if _, err := fmt.Sscanf(split[0], "%f", &v1); err != nil {
			panic(err)
		}
		if _, err := fmt.Sscanf(split[1], "%f", &v2); err != nil {
			panic(err)
		}

		if axC == 'y' {
			for i := v1; i <= v2; i++ {
				w[complex(c, i)] = -1
			}
			min = getMin(complex(c, v1), min)
		} else {
			for i := v1; i <= v2; i++ {
				w[complex(i, c)] = -1
			}
			max = getMax(complex(v2, c), max)
		}
	}

	return w, min, max
}

func getMin(p1, p2 complex64) complex64 {
	return complex(
		float32(math.Min(float64(real(p1)), float64(real(p2)))),
		float32(math.Min(float64(imag(p1)), float64(imag(p2)))))
}

func getMax(p1, p2 complex64) complex64 {
	return complex(
		float32(math.Max(float64(real(p1)), float64(real(p2)))),
		float32(math.Max(float64(imag(p1)), float64(imag(p2)))))
}

func less(c1, c2 complex64) bool {
	if real(c1) < real(c2) {
		return true
	}

	if real(c1) > real(c2) {
		return false
	}

	return imag(c1) < imag(c2)
}

func drawWorld(w world, min complex64, max complex64, filename string) {
	var width = int(imag(max)-imag(min)) + 5
	var height = int(real(max)-real(min)) + 5

	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for i := real(min); i <= real(max)+4; i++ {
		for j := imag(min) - 1; j <= imag(max)+4; j++ {
			var c color.RGBA
			switch w[complex(i, j)] {
			case 0:
				c = color.RGBA{0, 0, 0, 255}
			case 1:
				c = color.RGBA{0, 0, 255, 255}
			case -1:
				c = color.RGBA{255, 255, 255, 255}
			}
			im.SetRGBA(int(j-imag(min)+1), int(i-real(min)), c)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, im)
}
