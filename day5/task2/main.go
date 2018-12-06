package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
)

type result struct {
	r   rune
	l   int
	res string
}

func main() {
	var stdin io.Reader
	stdin = os.Stdin

	allRunes := `abcdefghijklmnopqrstuvwxyz`
	ins := make([]chan<- rune, 0, len(allRunes))
	outs := make([]<-chan result, 0, len(allRunes))
	for _, r := range allRunes {
		in, out := modCollapse(r)
		ins = append(ins, in)
		outs = append(outs, out)
	}

	r := bufio.NewReader(stdin)
	for {
		r, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}

		for _, in := range ins {
			in <- r
		}
	}
	for _, in := range ins {
		close(in)
	}

	var (
		min = math.MaxInt32
		res result
	)
	for _, out := range outs {
		l := <-out
		if l.l < min {
			min = l.l
			res = l
		}
	}
	fmt.Printf("%c: %s: %v\n", res.r, res.res, res.l)
}

func modCollapse(skip rune) (chan<- rune, <-chan result) {
	in := make(chan rune)
	out := make(chan result)
	skip = unicode.ToUpper(skip)

	go func() {
		var poly []rune
		for r := range in {
			if unicode.ToUpper(r) == skip {
				continue
			}

			if len(poly) > 0 && mustCollapse(poly[len(poly)-1], r) {
				poly = poly[:len(poly)-1]
				continue
			}

			poly = append(poly, r)
		}
		var resStr string
		for _, r := range poly {
			resStr += fmt.Sprintf("%c", r)
		}
		out <- result{skip, len(poly), resStr}
	}()

	return in, out
}

func mustCollapse(r1, r2 rune) bool {
	return xor(unicode.IsUpper(r1), unicode.IsUpper(r2)) && unicode.ToUpper(r1) == unicode.ToUpper(r2);
}

func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
