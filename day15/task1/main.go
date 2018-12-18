package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

const (
	power = 3
	hp    = 200
)

var (
	w     = make(map[complex64]byte)
	width = 0
	units = make(map[complex64]*unit)
	dirs  = []complex64{-1i, 1, 1i, -1,}
)

// x - im
// y - re

type unit struct {
	race  byte
	pos   complex64
	hp    int
	power int
}

func lessFn(slice []complex64) func(i, j int) bool {
	return func(i, j int) bool {
		return less(slice[i], slice[j])
	}
}

func main() {
	f, err := os.Open("day15/input-test.txt")
	if err != nil {
		panic(err)
	}

	w = readWorld(f)
	printWorld(w)

	var rounds int
loop:
	for {
		queue := makeQueue()

		for j, pos := range queue {
			if battleFinished() {
				break loop
			}
			_ = j
			u, ok := units[pos]
			if !ok {
				continue
			}

			if ok, enemy := shallAttach(u); ok {
				attack(enemy, u)
				continue
			}

			move(u)

			if ok, enemy := shallAttach(u); ok {
				attack(enemy, u)
				continue
			}
		}
		rounds++
		//fmt.Printf("After step #%d\n", rounds)
		//printWorld(w)
		//fmt.Println()
	}

	var hps int
	for _, u := range units {
		hps += u.hp
	}
	fmt.Println(rounds * hps)
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

func shallAttach(u *unit) (bool, *unit) {
	var res *unit
	for p, u1 := range units {
		if u1.race == u.race {
			continue
		}

		if d := u.pos - p; math.Abs(float64(real(d)))+math.Abs(float64(imag(d))) == 1 {
			if res == nil {
				res = u1
				continue
			}

			if u1.hp < res.hp {
				res = u1
				continue
			} else if u1.hp == res.hp && less(u1.pos, res.pos) {
				res = u1
				continue
			}
		}
	}

	return res != nil, res
}

func move(u *unit) bool {

	enemiesAround := findOpenEnemySquares(u.race)
	dir, ok := findDirection(u.pos, enemiesAround)
	if !ok {
		return false
	}

	delete(units, u.pos)
	w[u.pos] = '.'
	u.pos = u.pos + dir
	w[u.pos] = u.race
	units[u.pos] = u
	return true
}

func findDirection(from complex64, to []complex64) (complex64, bool) {

	if len(to) == 0 {
		return 0, false
	}

	targets := make(map[complex64]struct{})
	for _, p := range to {
		targets[p] = struct{}{}
	}

	visited := map[complex64]int{
		from: 0,
	}
	front := []complex64{from}

	var steps int
	var found []complex64
	for len(front) > 0 && len(found) == 0 {
		steps++
		startLen := len(front)

		for _, p := range front {
			for _, d := range dirs {
				if _, ok := visited[p+d]; !ok && w[p+d] == '.' {
					front = append(front, p+d)
					visited[p+d] = steps
				}

				if _, ok := targets[p+d]; ok {
					found = append(found, p+d)
				}
			}
		}
		front = front[startLen:]
	}

	if len(found) == 0 {
		return 0, false
	}

	sort.Slice(found, lessFn(found))
	front = []complex64{found[0]}
	for steps != 1 {
		startLen := len(front)

		for _, p := range front {
			for _, d := range dirs {
				if len, ok := visited[p+d]; ok && len == steps-1 {
					front = append(front, p+d)
				}
			}
		}
		front = front[startLen:]
		steps--
	}

	sort.Slice(front, lessFn(front))
	dir := front[0] - from
	return dir, true
}

func attack(enemy *unit, u *unit) {
	enemy.hp -= u.power
	if enemy.hp <= 0 {
		w[enemy.pos] = '.'
		delete(units, enemy.pos)
	}
}

func makeQueue() (res []complex64) {
	for pos, b := range w {
		if b == 'E' || b == 'G' {
			res = append(res, pos)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return less(res[i], res[j])
	})
	return res
}

func battleFinished() bool {
	var e, g bool
	for _, u := range units {
		if u.race == 'E' {
			e = true
		} else {
			g = true
		}
		if e && g {
			return false
		}
	}
	return true
}

func findOpenEnemySquares(unitRace byte) (res []complex64) {
	ks := []complex64{1, -1, 1i, -1i}
	for p, u := range units {
		if u.race == unitRace {
			continue
		}
		for _, k := range ks {
			if w[p+k] == '.' {
				res = append(res, p+k)
			}
		}
	}
	return res
}

func printWorld(w map[complex64]byte) {
	for x := 0; x < width; x++ {
		for y := 0; y < len(w)/width; y++ {
			fmt.Printf("%c", w[complex(float32(x), float32(y))])
		}
		fmt.Println()
	}
}

func readWorld(r io.Reader) map[complex64]byte {
	res := make(map[complex64]byte)
	s := bufio.NewScanner(r)
	ln := 0
	for s.Scan() {
		bts := s.Bytes()
		if len(bts) > 0 {
			width = len(bts)
		}

		for i, b := range bts {
			p := complex(float32(ln), float32(i))
			switch b {
			case 'E', 'G':
				units[p] = &unit{race: b, pos: p, hp: hp, power: power}
			}
			res[p] = b
		}
		ln++
	}
	return res
}
