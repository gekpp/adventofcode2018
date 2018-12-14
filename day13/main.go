package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
)

var (
	world map[complex64]byte
	cars  map[complex64]*car
)

type car struct {
	pos   complex64 //y=im, x=re
	dir   complex64
	turns complex64
	dead  bool
}

func (c *car) Tick() {
	delete(cars, c.pos)
	c.pos += c.dir
	if c1, ok := cars[c.pos]; ok {
		log.Printf("Bang: (%.0f,%.0f)", imag(c.pos), real(c.pos))
		delete(cars, c1.pos)
		c.dead = true
		c.dead = true
		return
	}

	switch world[c.pos] {
	case '+':
		c.dir = c.dir * c.turns
		if c.turns == -1i {
			c.turns = -1 * c.turns
		} else {
			c.turns = c.turns * -1i
		}
	case '/':
		c.dir = complex(- imag(c.dir), -real(c.dir))
	case '\\':
		c.dir = complex(imag(c.dir), real(c.dir))
	}
	cars[c.pos] = c
}

func main() {
	world, cars = readInput()

	for len(cars) > 1 {
		carPos := make([]*complex64, 0, len(cars))
		for _, c := range cars {
			carPos = append(carPos, &c.pos)
		}

		sort.Slice(carPos, func(i, j int) bool {
			if real(*carPos[i]) < real(*carPos[j]) {
				return true
			} else if real(*carPos[i]) > real(*carPos[j]) {
				return false
			} else {
				return imag(*carPos[i]) < imag(*carPos[j])
			}
		})

		for _, p := range carPos {
			c, ok := cars[*p]
			if !ok || c.dead {
				continue
			}
			cars[*p].Tick()
		}
	}

	for _, c := range cars {
		log.Printf("Last: (%.0f,%.0f)", imag(c.pos), real(c.pos))
		break
	}
}

func readInput() (world map[complex64]byte, cars map[complex64]*car) {
	r := bufio.NewReader(os.Stdin)
	var lnNum int
	dirs := map[byte]complex64{'>': 1i, '^': -1, '<': -1i, 'v': 1}
	world = make(map[complex64]byte)
	cars = make(map[complex64]*car)
	for {
		lnNum++
		bts, err := r.ReadBytes('\n')
		if err == io.EOF && len(bts) == 0 {
			break
		}

		for i, b := range bts {
			pos := complex(float32(lnNum-1), float32(i))
			switch b {
			case '>', '<', '^', 'v':
				cars[pos] = &car{pos: pos, turns: complex(0, 1), dir: dirs[b]}
			case '/', '\\', '+':
				world[pos] = b
			default:
				continue
			}
		}
	}
	return world, cars
}
