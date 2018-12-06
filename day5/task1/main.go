package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var stdin io.Reader
	stdin = os.Stdin

	r := bufio.NewReader(stdin)
	var poly []rune
	for {
		r, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}

		if len(poly) > 0 && mustSquash(poly[len(poly)-1], r) {
			poly = poly[:len(poly)-1]
			continue
		}

		poly = append(poly, r)
	}

	fmt.Println(len(poly))
	//for _, r := range poly {
	//	fmt.Printf("%c", r)
	//}
}

func mustSquash(r1, r2 rune) bool {
	return xor(unicode.IsUpper(r1), unicode.IsUpper(r2)) && unicode.ToUpper(r1) == unicode.ToUpper(r2);
}

func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
